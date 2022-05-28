#!/bin/bash

# run tests
python -m unittest discover "$CONTAINER_BASE_DIR"/tests || exit 1

# Analysing the code with pylint
pylint --max-line-length=120 --output-format=colorized \
    "$CONTAINER_BASE_DIR"/pipeline "$CONTAINER_BASE_DIR"/tests

# Lint with flake8
flake8 . --count --select=E9,F63,F7,F82 --show-source --statistics
flake8 . --count --exit-zero --max-complexity=10 --max-line-length=120 --statistics

# execute pipeline
python -u "$CONTAINER_BASE_DIR"/pipeline/pipeline.py
