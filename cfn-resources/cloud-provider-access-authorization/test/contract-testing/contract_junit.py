#!/usr/bin/env python3
# contract_junit.py
#
# Checks a pytest JUnit report produced by a contract-test run. The isolated
# contract-test scripts use it to distinguish a test failure from a fixture
# setup error, so the delete-create deviation is tolerated only in its exact
# expected form.
#
# Default mode: every selected case must pass, with no failures, errors or
# skips. With --expect-failure CASE: the named case must be the only case in
# the report and must fail.

import argparse
import sys
import xml.etree.ElementTree as ET


def read_cases(path):
    root = ET.parse(path).getroot()
    cases = []
    for testcase in root.iter("testcase"):
        if testcase.find("failure") is not None:
            outcome = "failed"
        elif testcase.find("error") is not None:
            outcome = "errored"
        elif testcase.find("skipped") is not None:
            outcome = "skipped"
        else:
            outcome = "passed"
        cases.append((testcase.get("name", ""), outcome))
    return cases


def main():
    parser = argparse.ArgumentParser()
    parser.add_argument("junit", help="pytest JUnit XML report")
    parser.add_argument(
        "--expect-failure",
        metavar="CASE",
        help="expect CASE to be the only case and to fail",
    )
    args = parser.parse_args()

    try:
        cases = read_cases(args.junit)
    except (ET.ParseError, OSError) as error:
        print(f"could not read {args.junit}: {error}")
        return 1

    if not cases:
        print("the JUnit report contains no test cases")
        return 1

    if args.expect_failure:
        others = [case for case in cases if case[0] != args.expect_failure]
        target = [case for case in cases if case[0] == args.expect_failure]
        if len(target) != 1 or target[0][1] != "failed" or others:
            print(f"expected only {args.expect_failure} to fail, got: {cases}")
            return 1
        print(f"{args.expect_failure} failed as expected")
        return 0

    bad = [case for case in cases if case[1] != "passed"]
    if bad:
        print(f"unexpected case outcomes: {bad}")
        return 1
    print(f"{len(cases)} case(s) passed")
    return 0


if __name__ == "__main__":
    sys.exit(main())
