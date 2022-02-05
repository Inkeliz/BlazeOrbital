#include "textflag.h"

TEXT ·onNextFrame(SB), NOSPLIT, $0
  CallImport
  RET

TEXT ·renderSatellites(SB), NOSPLIT, $0
  CallImport
  RET

TEXT ·renderSuns(SB), NOSPLIT, $0
  CallImport
  RET
