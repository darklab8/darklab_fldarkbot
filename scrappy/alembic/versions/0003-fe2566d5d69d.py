"""0003

Revision ID: fe2566d5d69d
Revises: d6b845375f36
Create Date: 2022-08-21 08:56:24.489422

"""
from alembic import op
import sqlalchemy as sa


# revision identifiers, used by Alembic.
revision = 'fe2566d5d69d'
down_revision = 'd6b845375f36'
branch_labels = None
depends_on = None


def upgrade() -> None:
    # ### commands auto generated by Alembic - please adjust! ###
    op.create_index(op.f('ix_players_name'), 'players', ['name'], unique=True)
    # ### end Alembic commands ###


def downgrade() -> None:
    # ### commands auto generated by Alembic - please adjust! ###
    op.drop_index(op.f('ix_players_name'), table_name='players')
    # ### end Alembic commands ###