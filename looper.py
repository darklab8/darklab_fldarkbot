"module for background tasks in the loop"
from discord.ext import tasks, commands
from channel import (
    delete_messages_older_than_n_seconds,
    handle_tagged_messages,
)


class Looper(commands.Cog):
    def __init__(self, bot, storage):
        self.index = 0
        self.bot = bot
        self.printer.start()
        self.storage = storage

    def cog_unload(self):
        self.printer.cancel()
        print('unloading')

    @tasks.loop(seconds=5.0)
    async def printer(self):

        print(self.index)
        self.index += 1

        data = self.storage.get_game_data()
        self.storage.save_channel_settings()
        channel_ids = [int(item) for item in self.storage.channels.keys()]
        print(channel_ids)
        for channel_id in channel_ids:
            pass
            # await delete_messages_older_than_n_seconds(self.bot,
            #                                            self.storage.unique_tag,
            #                                            10, channel_id)
            # await handle_tagged_messages(self.bot, self.storage.unique_tag,
            #                              channel_id)

    @printer.before_loop
    async def before_printer(self):
        print('waiting...')
        await self.bot.wait_until_ready()