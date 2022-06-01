from pipeline.config.config import DataBaseConfig
from pipeline.config.logger import get_logger
from pipeline.extract import extract_data
from pipeline.load import Loader
from pipeline.transform import Transformer


def run_pipeline() -> None:
    logger = get_logger('Pipeline')
    logger.debug('this is the pipeline writing from the logger')

    data = extract_data()
    new_cases = Transformer.get_new_cases_per_month_n_location(data)
    Loader(DataBaseConfig()).publish_record_list_to_db(new_cases)
