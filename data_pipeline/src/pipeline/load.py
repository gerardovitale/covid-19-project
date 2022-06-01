from datetime import datetime
from pathlib import Path
from typing import List, Dict, Any

from pandas import DataFrame

from pipeline.config.config import DataBaseConfig
from pipeline.config.logger import get_logger


class Loader:

    def __init__(self, database_config: DataBaseConfig):
        self.logger = get_logger('Loader')
        self.config = database_config

    def publish_record_list_to_db(self, record_list: List[Dict[str, Any]]) -> None:
        self.logger.debug('Publishing data into MongoDB Atlas')
        self.config.db.new_cases_per_month_n_location.insert_many(record_list)
        self.logger.debug('Data published')

    def save_dataframe_locally(self, file_name: str, df: DataFrame) -> None:
        filepath = Path('/app/data/{0}-{1}.csv'.format(
            datetime.now().strftime('%Y-%m-%d'), file_name))
        try:
            filepath.parent.mkdir(parents=True)
            df.to_csv(filepath, index=False)
        except FileExistsError as error:
            self.logger.error(error)
        finally:
            self.logger.debug('DataFrame saved in {0}'.format(filepath))

    def save_record_list(self, new_cases):
        pass
