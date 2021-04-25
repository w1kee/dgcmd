package dgcmd

import (
	"github.com/bwmarrin/discordgo"
)

type (
	Command struct {
		Names     []string
		Usage     string
		Validator func([]string) error
		Callback  func(s *discordgo.Session, m *discordgo.MessageCreate, args []string)
	}

	Handler struct {
		cmdMap map[string]*Command
		prefix string
	}

	HandleFunc = func(s *discordgo.Session, m *discordgo.MessageCreate)
)

func NewHandler(s *discordgo.Session) *Handler {
	h := &Handler{
		cmdMap: make(map[string]*Command),
		prefix: "!",
	}

	s.AddHandler(func(s *discordgo.Session, m *discordgo.MessageCreate) {
		naa := parseCommand(m.Content, h.prefix)
		if naa == nil {
			return
		}
		cmd, ok := h.cmdMap[naa[0]]
		if !ok {
			return
		}
		if cmd.Validator != nil {
			var err error
			if len(naa) > 1 {
				err = cmd.Validator(naa[1:])
			} else {
				err = cmd.Validator([]string{})
			}
			if err != nil {
				s.ChannelMessageSendEmbed(m.ChannelID, &discordgo.MessageEmbed{
					Title: "FAIL",
					Color: 0xff0033,
					Fields: []*discordgo.MessageEmbedField{
						{Name: "Error", Value: err.Error()},
						{Name: "Usage", Value: cmd.Usage},
					},
				})
				return
			}
		}

		cmd.Callback(s, m, naa[1:])
	})

	return h
}

func (h *Handler) Add(c *Command) error {
	for _, name := range c.Names {
		h.cmdMap[name] = c
	}
	return nil
}
