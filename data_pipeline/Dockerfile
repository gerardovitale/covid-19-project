ARG PY_IMAGE_VARIANT=slim
ARG PYTHON_VERSION=3.9.5

FROM python:${PYTHON_VERSION}-${PY_IMAGE_VARIANT} AS python3

RUN apt-get update && \
    python -m venv /opt/venv

ENV PATH=/opt/venv/bin:$PATH

COPY ./data_pipeline/requirements.txt ${CONTAINER_BASE_DIR}/requirements.txt

RUN pip install --upgrade pip && \
    pip install flake8 pylint && \
    pip install --no-cache-dir -r ${CONTAINER_BASE_DIR}/requirements.txt
    
FROM python:${PYTHON_VERSION}-${PY_IMAGE_VARIANT} AS build-image

ENV PYTHONUNBUFFERED 1

ARG CONTAINER_BASE_DIR=/app
ARG MODE
ARG DATA_URL
ARG MONGO_PASS

ENV CONTAINER_BASE_DIR=${CONTAINER_BASE_DIR} \
    MODE=${MODE} \
    DATA_URL=${DATA_URL} \
    MONGO_PASS=${MONGO_PASS} \
    PATH=/opt/venv/bin:$PATH

COPY --from=python3 /opt/venv /opt/venv

COPY ./data_pipeline/src ${CONTAINER_BASE_DIR}/
COPY ./data_pipeline/entrypoint.sh ${CONTAINER_BASE_DIR}/

WORKDIR ${CONTAINER_BASE_DIR}/
RUN chmod +x ./entrypoint.sh

ENTRYPOINT ["bash", "./entrypoint.sh", "${MODE}"]