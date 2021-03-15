package ginutil

func (g *GinWrap) PathParamOrDefault(key, def string) string {
	var val = g.Context.Param(key)
	if val == "" {
		return def
	} else {
		return val
	}
}

func (g *GinWrap) PathParamOrEmpty(key string) string {
	var val = g.Context.Param(key)
	if val == "" {
		return ""
	} else {
		return val
	}
}

func (g *GinWrap) QueryParamOrDefault(key, def string) string {
	val, ok := g.Context.GetQuery(key)
	if !ok {
		return def
	}
	if val == "" {
		return def
	}
	return val
}

func (g *GinWrap) QueryParamOrEmpty(key string) string {
	val, ok := g.Context.GetQuery(key)
	if !ok {
		return ""
	}
	if val == "" {
		return ""
	}
	return val
}

func (g *GinWrap) FormParamOrDefault(key, def string) string {
	val, ok := g.Context.GetPostForm(key)
	if !ok {
		return def
	}
	if val == "" {
		return def
	}
	return val
}
func (g *GinWrap) FormParamOrEmpty(key string) string {
	val, ok := g.Context.GetPostForm(key)
	if !ok {
		return ""
	}
	if val == "" {
		return ""
	}
	return val
}
