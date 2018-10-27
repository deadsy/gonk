//-----------------------------------------------------------------------------
/*

Font

*/
//-----------------------------------------------------------------------------

package nk

//-----------------------------------------------------------------------------

type FontAtlas struct {
}

func NewFontAtlas() *FontAtlas {
	return &FontAtlas{}
}

//NK_API void nk_font_atlas_init_default(struct nk_font_atlas*);
func (atlas *FontAtlas) InitDefault() {
}

//NK_API void nk_font_atlas_init(struct nk_font_atlas*, struct nk_allocator*);
func (atlas *FontAtlas) Init() {
}

// NK_API void nk_font_atlas_init_custom(struct nk_font_atlas*, struct nk_allocator *persistent, struct nk_allocator *transient);
func (atlas *FontAtlas) InitCustom() {
}

// NK_API void nk_font_atlas_begin(struct nk_font_atlas*);
func (atlas *FontAtlas) Begin() {
}

// NK_API struct nk_font *nk_font_atlas_add(struct nk_font_atlas*, const struct nk_font_config*);
func (atlas *FontAtlas) Add() {
}

// NK_API struct nk_font* nk_font_atlas_add_default(struct nk_font_atlas*, float height, const struct nk_font_config*);
func (atlas *FontAtlas) AddDefault() {
}

// NK_API struct nk_font* nk_font_atlas_add_from_memory(struct nk_font_atlas *atlas, void *memory, nk_size size, float height, const struct nk_font_config *config);
func (atlas *FontAtlas) AddFromMemory() {
}

// NK_API struct nk_font* nk_font_atlas_add_from_file(struct nk_font_atlas *atlas, const char *file_path, float height, const struct nk_font_config*);
func (atlas *FontAtlas) AddFromFile() {
}

// NK_API struct nk_font *nk_font_atlas_add_compressed(struct nk_font_atlas*, void *memory, nk_size size, float height, const struct nk_font_config*);
func (atlas *FontAtlas) AddCompressed() {
}

// NK_API struct nk_font* nk_font_atlas_add_compressed_base85(struct nk_font_atlas*, const char *data, float height, const struct nk_font_config *config);
func (atlas *FontAtlas) AddCompressedBase85() {
}

// NK_API const void* nk_font_atlas_bake(struct nk_font_atlas*, int *width, int *height, enum nk_font_atlas_format);
func (atlas *FontAtlas) Bake() {
}

// NK_API void nk_font_atlas_end(struct nk_font_atlas*, nk_handle tex, struct nk_draw_null_texture*);
func (atlas *FontAtlas) End() {
}

// NK_API void nk_font_atlas_cleanup(struct nk_font_atlas *atlas);
func (atlas *FontAtlas) Cleanup() {
}

// NK_API void nk_font_atlas_clear(struct nk_font_atlas*);
func (atlas *FontAtlas) Clear() {
}

//-----------------------------------------------------------------------------
