#!/bin/bash

# run tests
python -m unittest discover $CONTAINER_BASE_DIR/tests || exit 1

# execute pipeline
python -u $CONTAINER_BASE_DIR/pipeline/pipeline.py
