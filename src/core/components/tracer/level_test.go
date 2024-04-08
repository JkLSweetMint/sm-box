package tracer

import (
	"testing"
)

func TestLevel_String(t *testing.T) {
	tests := []struct {
		name  string
		level Level
		want  string
	}{
		{
			name:  "Case 1",
			level: LevelMain,
			want:  allLevelsString[LevelMain-1],
		},
		{
			name:  "Case 2",
			level: LevelUseCaseEvent,
			want:  allLevelsString[LevelUseCaseEvent-1],
		},
		{
			name:  "Case 3",
			level: LevelRepositoryEvent,
			want:  allLevelsString[LevelRepositoryEvent-1],
		},
		{
			name:  "Case 4",
			level: LevelUseCaseInternal,
			want:  allLevelsString[LevelUseCaseInternal-1],
		},
		{
			name:  "Case 5",
			level: LevelPackageDebug,
			want:  allLevelsString[LevelPackageDebug-1],
		},
		{
			name:  "Case 6",
			level: LevelCoreEvent,
			want:  allLevelsString[LevelCoreEvent-1],
		},
		{
			name:  "Case 7",
			level: LevelComponentEvent,
			want:  allLevelsString[LevelComponentEvent-1],
		},
		{
			name:  "Case 8",
			level: LevelCoreAddonInternal,
			want:  allLevelsString[LevelCoreAddonInternal-1],
		},
		{
			name:  "Case 9",
			level: LevelCore,
			want:  allLevelsString[LevelCore-1],
		},
		{
			name:  "Case 10",
			level: LevelCoreTransportInternal,
			want:  allLevelsString[LevelCoreTransportInternal-1],
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.level.String(); got != tt.want {
				t.Errorf("String() = %v, want %v", got, tt.want)
			}
		})
	}
}
