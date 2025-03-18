import json
import logging
import requests # this requires adding request layer to lambda function
from urllib.parse import urlparse, urlunparse

logging.basicConfig(level=logging.DEBUG)
logger = logging.getLogger()

EC2_PROXY_ENDPOINT = "http://XX.X.X.XX" # Replace with private IP of EC2 proxy running Python/Flask
if not EC2_PROXY_ENDPOINT:
    logger.error("EC2_PROXY_ENDPOINT is not set")
    raise ValueError("EC2_PROXY_ENDPOINT is not set")
logger.debug(f"EC2_PROXY_ENDPOINT: {EC2_PROXY_ENDPOINT}")

def lambda_handler(event, context):
    """
    Expected event example:
    {
        "method": "GET",
        "url": "https://cloud-dev.mongodb.com/api/atlas/v2/groups",
        "headers": {"Header-Key": "value", ...},
        "body": "request body as string"
    }

    This Lambda function should be deployed in a private subnet. It's corresponding SG 
    only allows traffic to EC2 proxy running Python/Flask
    """
    logger.debug(f"Received event: {json.dumps(event)}") # TODO: remove
    try:
        method = event.get("method")
        url = event.get("url")
        headers = event.get("headers", {})
        body = event.get("body", None)

        if method is None or url is None:
            msg = "Missing 'method' or 'url' in the event payload"
            logger.error(msg)
            raise ValueError(msg)

        logger.debug(f"Forwarding request - Method: {method}, URL: {url}, Headers: {headers}, Body: {body}")

        parsed_url = urlparse(url)
        if parsed_url.scheme and parsed_url.netloc:
            # Extract the path and query string because incoming URL will be in format:
            # https://www.cloud.mongodb.com/api/atlas/v2/groups?param=value
            path = parsed_url.path or ""
            query = f"?{parsed_url.query}" if parsed_url.query else ""
            new_url = path + query
            logger.debug(f"Extracted relative URL: {new_url} from absolute URL: {url}")
        else:
            new_url = url
            logger.debug(f"URL is relative: {new_url}")

        # Construct URL for the EC2 proxy
        full_url = EC2_PROXY_ENDPOINT.rstrip("/") + new_url
        logger.debug(f"Constructed full URL for proxy: {full_url}")

        # Forward request to the EC2 proxy
        response = requests.request(method, full_url, headers=headers, data=body)
        logger.debug(f"Response from EC2 proxy - Status: {response.status_code}, Headers: {dict(response.headers)}, Body: {response.text}")

        return {
            "statusCode": response.status_code,
            "headers": dict(response.headers),
            "body": response.text
        }
    except Exception as e:
        logger.exception("Error processing the request:")
        return {
            "statusCode": 500,
            "headers": {},
            "body": json.dumps({"error": str(e)})
        }
