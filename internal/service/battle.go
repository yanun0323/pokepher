package service

import (
	"main/internal/domain/entity"
	"main/internal/utils"
)

type BattleStage string

const (
	_battleStageEntry  BattleStage = "entry"
	_battleStageSelect BattleStage = "select"
	_battleStageWin    BattleStage = "win"
	_battleStageLose   BattleStage = "lose"
)

type Battle struct {
	BattleStage *utils.Value[BattleStage]

	Player   *entity.Character
	Opponent *entity.Character

	PlayerPokemon   []*entity.Pokemon
	OpponentPokemon []*entity.Pokemon
}

func NewBattle(Player, Opponent *entity.Character, PlayerPokemon, OpponentPokemon []*entity.Pokemon) *Battle {
	return &Battle{
		BattleStage: utils.NewValue(_battleStageEntry),
	}
}
