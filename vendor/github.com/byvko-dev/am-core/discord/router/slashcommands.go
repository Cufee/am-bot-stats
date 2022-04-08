package router

import (
	"context"
	"fmt"
	"strings"

	"github.com/andersfylling/disgord"
	"github.com/byvko-dev/am-core/helpers/slices"
	"github.com/byvko-dev/am-core/logs"
)

func RegisterSlashCommands(client *disgord.Client) {
	var commands []disgord.CreateApplicationCommand

	for _, cmd := range validCommands {
		if cmd.AddSlashCommand {
			// Check if a command or alias is already registered
			if slices.Contains(boundCommands, fmt.Sprintf("%v-slash", cmd.Name)) > -1 {
				logs.Error(fmt.Sprintf("Command %v is already bound", fmt.Sprintf("%v-slash", cmd.Name)))
				continue
			}
			for _, a := range cmd.Aliases {
				if slices.Contains(boundCommands, fmt.Sprintf("%v-slash", a)) > -1 {
					logs.Error(fmt.Sprintf("Alias %v is already bound", fmt.Sprintf("%v-slash", a)))
					continue
				}
				boundCommands = append(boundCommands, fmt.Sprintf("%v-slash", a))
			}
			boundCommands = append(boundCommands, fmt.Sprintf("%v-slash", cmd.Name))

			var args []*disgord.ApplicationCommandOption
			for _, a := range cmd.Arguments {
				arg := disgord.ApplicationCommandOption(a)
				arg.Type = disgord.OptionTypeString
				arg.Autocomplete = len(arg.Choices) > 0
				args = append(args, &arg)
			}

			var command = disgord.CreateApplicationCommand{
				Name:              strings.ToLower(cmd.Name),
				DefaultPermission: true,
				Description:       cmd.Description,
				Options:           args,
				Type:              disgord.ApplicationCommandChatInput,
			}
			commands = append(commands, command)
		}
	}

	// Get all registered commands
	currentCommandsPointer, err := client.ApplicationCommand(733715379712557098).Global().List()
	if err != nil {
		logs.Error("Failed to list commands", err)
	}
	currentCommands := *currentCommandsPointer
	currentNames := make([]string, len(currentCommands))
	for i, c := range currentCommands {
		currentNames[i] = strings.ToLower(c.Name)
	}
	logs.Info("Current commands: %v", currentNames)

	// Create new commands and update existing ones
	for _, cmd := range commands {
		cmdIndex := slices.Contains(currentNames, strings.ToLower(cmd.Name))
		if cmdIndex == -1 {
			err := client.ApplicationCommand(733715379712557098).Global().Create(&cmd)
			if err != nil {
				logs.Error("Failed to create command %v: %v", cmd.Name, err)
			}
		} else {
			var update = disgord.UpdateApplicationCommand{
				Name:              &cmd.Name,
				DefaultPermission: &cmd.DefaultPermission,
				Description:       &cmd.Description,
				Options:           &cmd.Options,
			}
			err = client.ApplicationCommand(733715379712557098).Global().Update(currentCommands[cmdIndex].ID, &update)
			if err != nil {
				logs.Error("Error updating slash command for %v: %v", cmd.Name, err)
				logs.Error("%+v", currentCommands[cmdIndex])
			}
		}
	}

	// Delete commands that are no longer registered
	newCommands := make([]string, len(commands))
	for i, c := range commands {
		newCommands[i] = strings.ToLower(c.Name)
	}
	for _, cmd := range currentCommands {
		if slices.Contains(newCommands, strings.ToLower(cmd.Name)) == -1 {
			err := client.ApplicationCommand(733715379712557098).Global().Delete(cmd.ID)
			if err != nil {
				logs.Error("Error deleting slash command for %v: %v", cmd.Name, err)
			}
		}
	}
}

func AddSlashCommandsHandlers(options RouterOptions) func(client *disgord.Client) {
	options.check() // Panic if some required handlers are missing

	return func(client *disgord.Client) {
		client.Gateway().InteractionCreate(func(s disgord.Session, h *disgord.InteractionCreate) {
			for _, cmd := range validCommands {
				if cmd.AddSlashCommand && h.Data != nil && h.Data.Name == cmd.Name {
					replyToInteraction(s, h, "<a:ExpectDelays:862151171581542470>")

					// Make a list of compatible arguments
					arguments := make([]string, len(h.Data.Options))
					for i, o := range h.Data.Options {
						arguments[i] = fmt.Sprint(o.Value)
					}

					// We are not able to reply from within the handler, so we reply with a loading message and then edit it
					var response disgord.UpdateMessage
					runCommandHandler(Command{
						Name:         cmd.Name,
						Arguments:    arguments,
						session:      s,
						options:      cmd,
						userID:       h.Member.UserID.String(),
						message:      h.Message, // Message here is always nil
						userCheck:    options.UserCheckHandler,
						reportError:  options.ErrorReportHandler,
						errorPrinter: options.ErrorPrinter,
						Reply: func(c Command, content ...interface{}) error {
							for _, c := range content {
								switch chunk := c.(type) {
								case disgord.CreateMessageFile: // File
									response.File = &chunk

								default: // Assume it is a string
									chunks := []string{fmt.Sprint(chunk)}
									if response.Content != nil {
										chunks = append([]string{*response.Content}, chunks...)
									}
									content := strings.Join(chunks, "\n")
									response.Content = &content
								}
							}
							return nil
						},
					})

					// It is possible to call Reply() multiple times, so we need to join the response
					defer func() {
						s.EditInteractionResponse(context.Background(), h, &response)
					}()
					return // Allows to catch unknown commands
				}
			}
			// If a command is not found, it was likely registered in a different module
		})
	}
}

func replyToInteraction(s disgord.Session, h *disgord.InteractionCreate, content string) {
	s.SendInteractionResponse(context.Background(), h, &disgord.CreateInteractionResponse{
		Type: 4,
		Data: &disgord.CreateInteractionResponseData{
			Content: content,
		},
	})
}
