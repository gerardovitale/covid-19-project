import os

import pandas as pd

from pipeline.config.logger import get_logger

logger = get_logger('Extractor')


def extract_data(chunk: bool = False) -> pd.DataFrame:
    chunksize = 1000 if chunk else None
    logger.debug('Extracting data')
    return pd.read_csv(os.getenv('DATA_URL'), nrows=chunksize)
