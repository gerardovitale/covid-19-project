from pipeline.config.logger import get_logger


def run_pipeline() -> None:
    logger = get_logger('Pipeline')
    logger.info('this is the pipeline writing from the logger')
