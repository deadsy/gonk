//-----------------------------------------------------------------------------
/*

Common Code

*/
//-----------------------------------------------------------------------------

package nk

//-----------------------------------------------------------------------------

// Color is a color expressed with uint8 component values.
type Color struct {
	r, g, b, a uint8
}

// Colorf is a color expressed with float32 component values.
type Colorf struct {
	r, g, b, a float32
}

// Vec2 is a 2d vector with float32 values.
type Vec2 struct {
	x, y float32
}

// Vec2i is a 2d vector with uint16 values.
type Vec2i struct {
	x, y uint16
}

// Rect is a rectangle with float32 values.
type Rect struct {
	x, y, w, h float32
}

// Recti is a rectangle with uint16 values.
type Recti struct {
	x, y, w, h uint16
}

//-----------------------------------------------------------------------------

type Handle struct {
	ptr interface{}
	id  int
}

//-----------------------------------------------------------------------------

type mouse struct {
	pos         Vec2
	prev        Vec2
	delta       Vec2
	scrollDelta Vec2
	grab        bool
	grabbed     bool
	ungrab      bool
}

type keyboard struct {
}

// Input records the current input state.
type Input struct {
	keyboard keyboard
	mouse    mouse
}

//-----------------------------------------------------------------------------

// Context is the top-level data structure for the UI.
type Context struct {

	// public
	input Input
	//struct nk_style style;
	//struct nk_buffer memory;
	//struct nk_clipboard clip;
	//nk_flags last_widget_state;
	//enum nk_button_behavior button_behavior;
	//struct nk_configuration_stacks stacks;
	//float delta_time_seconds;

}

//-----------------------------------------------------------------------------
