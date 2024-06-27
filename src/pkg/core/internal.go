package core

import "sm-box/pkg/core/components/tracer"

// updateState - обновление состояния ядра.
func (c *core) updateState(state State) (err error) {
	// tracer
	{
		var trc = tracer.New(tracer.LevelCoreInternal)

		trc.FunctionCall(state)
		defer func() { trc.Error(err).FunctionCallFinished() }()
	}

	var old = c.State()

	defer func() {
		if err != nil {
			c.Components().Logger().Error().
				Format("An error occurred during the core status update: '%s'. ", err).
				Field("old_state", old).
				Field("new_state", state).Write()
			return
		}

		c.Components().Logger().Info().
			Format("The state of the system core has been changed from '%s' to '%s'. ",
				old,
				instance.State(),
			).Write()
	}()

	switch state {
	case StateNew:
		{
			if old == StateNil {
				c.state = &stateNew{
					components: c.components,
					tools:      c.tools,

					ctx:  c.ctx,
					conf: c.conf,
				}
				return
			}
		}
	case StateBooted:
		{
			if old == StateNew {
				c.state = &stateBooted{
					components: c.components,
					tools:      c.tools,

					ctx:  c.ctx,
					conf: c.conf,
				}
				return
			}
		}
	case StateServed:
		{
			if old == StateBooted {
				c.state = &stateServed{
					components: c.components,
					tools:      c.tools,

					ctx:  c.ctx,
					conf: c.conf,
				}
				return
			}
		}
	case StateOff:
		{
			if old == StateServed {
				c.state = &stateOff{
					components: c.components,
					tools:      c.tools,

					ctx:  c.ctx,
					conf: c.conf,
				}
				return
			}
		}
	}

	err = ErrInvalidSystemCoreState
	return
}
