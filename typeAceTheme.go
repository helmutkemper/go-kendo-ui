package telerik

type AceTheme int

var aceThemes = [...]string {
  "",
  "ace/theme/chrome",
  "ace/theme/clouds",
  "ace/theme/crimson_editor",
  "ace/theme/dawn",
  "ace/theme/dreamweaver",
  "ace/theme/eclipse",
  "ace/theme/github",
  "ace/theme/iplastic",
  "ace/theme/solarized_light",
  "ace/theme/textmate",
  "ace/theme/tomorrow",
  "ace/theme/xcode",
  "ace/theme/kuroir",
  "ace/theme/katzenmilch",
  "ace/theme/sqlserver",
  "ace/theme/ambiance",
  "ace/theme/chaos",
  "ace/theme/clouds_midnight",
  "ace/theme/dracula",
  "ace/theme/cobalt",
  "ace/theme/gruvbox",
  "ace/theme/gob",
  "ace/theme/idle_fingers",
  "ace/theme/kr_theme",
  "ace/theme/merbivore",
  "ace/theme/merbivore_soft",
  "ace/theme/mono_industrial",
  "ace/theme/monokai",
  "ace/theme/pastel_on_dark",
  "ace/theme/solarized_dark",
  "ace/theme/terminal",
  "ace/theme/tomorrow_night",
  "ace/theme/tomorrow_night_blue",
  "ace/theme/tomorrow_night_bright",
  "ace/theme/tomorrow_night_eighties",
  "ace/theme/twilight",
  "ace/theme/vibrant_ink",
}

func (el AceTheme) String() string {
  return aceThemes[el]
}

const (
  ACE_THEME_CHROME AceTheme = iota + 1
  ACE_THEME_CLOUDS
  ACE_THEME_CRIMSON_EDITOR
  ACE_THEME_DAWN
  ACE_THEME_DREAMWEAVER
  ACE_THEME_ECLIPSE
  ACE_THEME_GITHUB
  ACE_THEME_IPLASTIC
  ACE_THEME_SOLARIZED_LIGHT
  ACE_THEME_TEXTMATE
  ACE_THEME_TOMORROW
  ACE_THEME_XCODE
  ACE_THEME_KUROIR
  ACE_THEME_KATZENMILCH
  ACE_THEME_SQLSERVER
  ACE_THEME_AMBIANCE
  ACE_THEME_CHAOS
  ACE_THEME_CLOUDS_MIDNIGHT
  ACE_THEME_DRACULA
  ACE_THEME_COBALT
  ACE_THEME_GRUVBOX
  ACE_THEME_GOB
  ACE_THEME_IDLE_FINGERS
  ACE_THEME_KR_THEME
  ACE_THEME_MERBIVORE
  ACE_THEME_MERBIVORE_SOFT
  ACE_THEME_MONO_INDUSTRIAL
  ACE_THEME_MONOKAI
  ACE_THEME_PASTEL_ON_DARK
  ACE_THEME_SOLARIZED_DARK
  ACE_THEME_TERMINAL
  ACE_THEME_TOMORROW_NIGHT
  ACE_THEME_TOMORROW_NIGHT_BLUE
  ACE_THEME_TOMORROW_NIGHT_BRIGHT
  ACE_THEME_TOMORROW_NIGHT_EIGHTIES
  ACE_THEME_TWILIGHT
  ACE_THEME_VIBRANT_INK
)
