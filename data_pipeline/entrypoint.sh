#!/bin/bash

# run tests
echo "***************************************************************************"
echo "[INFO] running python unittest"
python -m unittest discover "$CONTAINER_BASE_DIR"/ || exit 1

if [[ $MODE == 'lint' ]]; then
  # Analysing the code with pylint
  echo "***************************************************************************"
  echo "[INFO] running pylint"
  pylint --max-line-length=120 --output-format=colorized \
      "$CONTAINER_BASE_DIR"/pipeline "$CONTAINER_BASE_DIR"/tests

  # Lint with flake8
  echo "***************************************************************************"
  echo "[INFO] running flake8"
  flake8 . --count --select=E9,F63,F7,F82 --show-source --statistics
  flake8 . --count --exit-zero --max-complexity=10 --max-line-length=120 --statistics
fi

if [[ $MODE == 'pipe' ]]; then
  # execute pipeline
  echo "***************************************************************************"
  echo "[INFO] executing data pipeline"
  python -u "$CONTAINER_BASE_DIR"/run.py
fi
