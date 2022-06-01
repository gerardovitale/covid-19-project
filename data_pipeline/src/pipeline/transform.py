from datetime import datetime
from typing import Any, Dict, List

from pandas import DataFrame

from pipeline.config.logger import get_logger


class Transformer:
    logger = get_logger('Transformer')

    @classmethod
    def get_new_cases_per_month_n_location(cls, data: DataFrame) -> List[Dict[str, Any]]:
        preprocessed_data = cls._perform_common_transformation(data)
        cls.logger.debug('Commons transformation done')
        columns_to_include = ['date', 'location', 'continent', 'total_cases', 'new_cases']
        new_cases = preprocessed_data[columns_to_include].copy()
        new_cases.set_index('date', inplace=True)
        new_cases['month'] = new_cases.index.month
        new_cases['year'] = new_cases.index.year
        new_cases_per_month_n_country = new_cases.groupby(by=['year', 'month', 'location']).sum()
        new_cases_per_month_n_country.reset_index(inplace=True)
        cls.logger.debug('returning new_cases_per_month_n_location as List[Dict]')
        return new_cases_per_month_n_country.to_dict('records')

    @staticmethod
    def _perform_common_transformation(data: DataFrame):
        data.dropna(subset=['continent'], inplace=True)
        data['date'] = data['date'].apply(lambda date: datetime.strptime(date, '%Y-%m-%d'))
        # data = self._add_missing_january_records(data)
        return data

    def _add_missing_january_records(self):
        pass
