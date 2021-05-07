"module building all data to be stored"
import os
from types import SimpleNamespace
import json
from dotenv import load_dotenv
import requests


class Storage():
    def __init__(self, unique_tag='dark_info:'):
        self.unique_tag = unique_tag
        self.settings = self.load_env_settings()
        self.channels = self.load_channel_settings()

    def load_env_settings(self) -> SimpleNamespace:
        "loading settings from os environment"
        load_dotenv()

        output = SimpleNamespace()
        for item, value in os.environ.items():
            setattr(output, item, value)
        return output

    def load_channel_settings(self) -> dict:
        """loadding perssistent settings
        set by users about channels"""
        output = {}
        try:
            with open('channels.json', 'r') as file_:
                output = json.loads(file_.read())
        except FileNotFoundError:
            print('ERR failed to load channels.json')
        return output

    def save_channel_settings(self) -> None:
        """loadding perssistent settings
        set by users about channels"""
        try:
            with open('channels.json', 'w') as file_:
                file_.write(json.dumps(self.channels, indent=2))
        except OSError:
            print('ERR failed to save channels.json')

    def get_game_data(self) -> SimpleNamespace:
        output = SimpleNamespace()
        output.players = requests.get(self.settings.player_request_url).json()
        output.bases = requests.get(self.settings.base_request_url).json()
        return output

    def base_add(self, name):
        print('adding the base')

    # def get_channel_data(self, key) -> SimpleNamespace:
    #     return deepcopy(self.storage.channels[key])