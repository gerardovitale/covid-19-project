import os
import pandas as pd


def extract_data(chunk: bool = False) -> pd.DataFrame:
    DATA_URL = os.getenv('DATA_URL')
    chunksize = 1000 if chunk else None 
    data = pd.read_csv(DATA_URL, nrows=chunksize)
    return data
