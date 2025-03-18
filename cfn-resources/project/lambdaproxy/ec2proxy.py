from flask import Flask, request, Response
import requests
import logging

app = Flask(__name__)

logging.basicConfig(level=logging.DEBUG)
logger = logging.getLogger(__name__)

# NOTE: additional configuration would be required to also support Realm
TARGET_SERVER = "https://cloud-dev.mongodb.com"
logger.debug(f"EC2 Proxy configured with TARGET_SERVER: {TARGET_SERVER}")

@app.route('/', defaults={'path': ''}, methods=["GET", "POST", "PUT", "DELETE", "PATCH", "OPTIONS", "HEAD"])
@app.route('/<path:path>', methods=["GET", "POST", "PUT", "DELETE", "PATCH", "OPTIONS", "HEAD"])
def proxy(path):
    logger.debug(f"Received request for path: {path}")
    # Build the target URL
    target_url = f"{TARGET_SERVER}/{path}"
    logger.debug(f"Target URL: {target_url}")
    
    # Copy the incoming headers
    headers = {key: value for key, value in request.headers if key.lower() != 'host'}
    logger.debug(f"Request headers: {headers}")
    
    # Forward the request to Atlas
    try:
        resp = requests.request(
            method=request.method,
            url=target_url,
            headers=headers,
            data=request.get_data(),
            cookies=request.cookies,
            allow_redirects=False
        )
        # logger.debug(f"Received response from target server - Status: {resp.status_code}, Headers: {resp.headers}")
    except Exception as e:
        logger.exception("Error forwarding the request to the target server:")
        return Response("Error forwarding request", status=500)
    
    excluded_headers = ['content-encoding', 'content-length', 'transfer-encoding', 'connection']
    response_headers = [(name, value) for (name, value) in resp.raw.headers.items()
                        if name.lower() not in excluded_headers]
    
    response = Response(resp.content, resp.status_code, response_headers)
    logger.debug(f"Returning proxied response with status: {resp.status_code}")
    return response

if __name__ == '__main__':
    logger.debug("Starting EC2 Proxy Flask app on 0.0.0.0:80")
    app.run(host='0.0.0.0', port=80)
