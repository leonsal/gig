package imgui

var colorMap = map[string]Vec4{
	"Aliceblue":            Vec4{0.941, 0.973, 1.000, 1.0},
	"Antiquewhite":         Vec4{0.980, 0.922, 0.843, 1.0},
	"Aqua":                 Vec4{0.000, 1.000, 1.000, 1.0},
	"Aquamarine":           Vec4{0.498, 1.000, 0.831, 1.0},
	"Azure":                Vec4{0.941, 1.000, 1.000, 1.0},
	"Beige":                Vec4{0.961, 0.961, 0.863, 1.0},
	"Bisque":               Vec4{1.000, 0.894, 0.769, 1.0},
	"Black":                Vec4{0.000, 0.000, 0.000, 1.0},
	"Blanchedalmond":       Vec4{1.000, 0.922, 0.804, 1.0},
	"Blue":                 Vec4{0.000, 0.000, 1.000, 1.0},
	"Blueviolet":           Vec4{0.541, 0.169, 0.886, 1.0},
	"Brown":                Vec4{0.647, 0.165, 0.165, 1.0},
	"Burlywood":            Vec4{0.871, 0.722, 0.529, 1.0},
	"Cadetblue":            Vec4{0.373, 0.620, 0.627, 1.0},
	"Chartreuse":           Vec4{0.498, 1.000, 0.000, 1.0},
	"Chocolate":            Vec4{0.824, 0.412, 0.118, 1.0},
	"Coral":                Vec4{1.000, 0.498, 0.314, 1.0},
	"Cornflowerblue":       Vec4{0.392, 0.584, 0.929, 1.0},
	"Cornsilk":             Vec4{1.000, 0.973, 0.863, 1.0},
	"Crimson":              Vec4{0.863, 0.078, 0.235, 1.0},
	"Cyan":                 Vec4{0.000, 1.000, 1.000, 1.0},
	"Darkblue":             Vec4{0.000, 0.000, 0.545, 1.0},
	"Darkcyan":             Vec4{0.000, 0.545, 0.545, 1.0},
	"Darkgoldenrod":        Vec4{0.722, 0.525, 0.043, 1.0},
	"Darkgray":             Vec4{0.663, 0.663, 0.663, 1.0},
	"Darkgreen":            Vec4{0.000, 0.392, 0.000, 1.0},
	"Darkgrey":             Vec4{0.663, 0.663, 0.663, 1.0},
	"Darkkhaki":            Vec4{0.741, 0.718, 0.420, 1.0},
	"Darkmagenta":          Vec4{0.545, 0.000, 0.545, 1.0},
	"Darkolivegreen":       Vec4{0.333, 0.420, 0.184, 1.0},
	"Darkorange":           Vec4{1.000, 0.549, 0.000, 1.0},
	"Darkorchid":           Vec4{0.600, 0.196, 0.800, 1.0},
	"Darkred":              Vec4{0.545, 0.000, 0.000, 1.0},
	"Darksalmon":           Vec4{0.914, 0.588, 0.478, 1.0},
	"Darkseagreen":         Vec4{0.561, 0.737, 0.561, 1.0},
	"Darkslateblue":        Vec4{0.282, 0.239, 0.545, 1.0},
	"Darkslategray":        Vec4{0.184, 0.310, 0.310, 1.0},
	"Darkslategrey":        Vec4{0.184, 0.310, 0.310, 1.0},
	"Darkturquoise":        Vec4{0.000, 0.808, 0.820, 1.0},
	"Darkviolet":           Vec4{0.580, 0.000, 0.827, 1.0},
	"Deeppink":             Vec4{1.000, 0.078, 0.576, 1.0},
	"Deepskyblue":          Vec4{0.000, 0.749, 1.000, 1.0},
	"Dimgray":              Vec4{0.412, 0.412, 0.412, 1.0},
	"Dimgrey":              Vec4{0.412, 0.412, 0.412, 1.0},
	"Dodgerblue":           Vec4{0.118, 0.565, 1.000, 1.0},
	"Firebrick":            Vec4{0.698, 0.133, 0.133, 1.0},
	"Floralwhite":          Vec4{1.000, 0.980, 0.941, 1.0},
	"Forestgreen":          Vec4{0.133, 0.545, 0.133, 1.0},
	"Fuchsia":              Vec4{1.000, 0.000, 1.000, 1.0},
	"Gainsboro":            Vec4{0.863, 0.863, 0.863, 1.0},
	"Ghostwhite":           Vec4{0.973, 0.973, 1.000, 1.0},
	"Gold":                 Vec4{1.000, 0.843, 0.000, 1.0},
	"Goldenrod":            Vec4{0.855, 0.647, 0.125, 1.0},
	"Gray":                 Vec4{0.502, 0.502, 0.502, 1.0},
	"Green":                Vec4{0.000, 0.502, 0.000, 1.0},
	"Greenyellow":          Vec4{0.678, 1.000, 0.184, 1.0},
	"Grey":                 Vec4{0.502, 0.502, 0.502, 1.0},
	"Honeydew":             Vec4{0.941, 1.000, 0.941, 1.0},
	"Hotpink":              Vec4{1.000, 0.412, 0.706, 1.0},
	"Indianred":            Vec4{0.804, 0.361, 0.361, 1.0},
	"Indigo":               Vec4{0.294, 0.000, 0.510, 1.0},
	"Ivory":                Vec4{1.000, 1.000, 0.941, 1.0},
	"Khaki":                Vec4{0.941, 0.902, 0.549, 1.0},
	"Lavender":             Vec4{0.902, 0.902, 0.980, 1.0},
	"Lavenderblush":        Vec4{1.000, 0.941, 0.961, 1.0},
	"Lawngreen":            Vec4{0.486, 0.988, 0.000, 1.0},
	"Lemonchiffon":         Vec4{1.000, 0.980, 0.804, 1.0},
	"Lightblue":            Vec4{0.678, 0.847, 0.902, 1.0},
	"Lightcoral":           Vec4{0.941, 0.502, 0.502, 1.0},
	"Lightcyan":            Vec4{0.878, 1.000, 1.000, 1.0},
	"Lightgoldenrodyellow": Vec4{0.980, 0.980, 0.824, 1.0},
	"Lightgray":            Vec4{0.827, 0.827, 0.827, 1.0},
	"Lightgreen":           Vec4{0.565, 0.933, 0.565, 1.0},
	"Lightgrey":            Vec4{0.827, 0.827, 0.827, 1.0},
	"Lightpink":            Vec4{1.000, 0.714, 0.757, 1.0},
	"Lightsalmon":          Vec4{1.000, 0.627, 0.478, 1.0},
	"Lightseagreen":        Vec4{0.125, 0.698, 0.667, 1.0},
	"Lightskyblue":         Vec4{0.529, 0.808, 0.980, 1.0},
	"Lightslategray":       Vec4{0.467, 0.533, 0.600, 1.0},
	"Lightslategrey":       Vec4{0.467, 0.533, 0.600, 1.0},
	"Lightsteelblue":       Vec4{0.690, 0.769, 0.871, 1.0},
	"Lightyellow":          Vec4{1.000, 1.000, 0.878, 1.0},
	"Lime":                 Vec4{0.000, 1.000, 0.000, 1.0},
	"Limegreen":            Vec4{0.196, 0.804, 0.196, 1.0},
	"Linen":                Vec4{0.980, 0.941, 0.902, 1.0},
	"Magenta":              Vec4{1.000, 0.000, 1.000, 1.0},
	"Maroon":               Vec4{0.502, 0.000, 0.000, 1.0},
	"Mediumaquamarine":     Vec4{0.400, 0.804, 0.667, 1.0},
	"Mediumblue":           Vec4{0.000, 0.000, 0.804, 1.0},
	"Mediumorchid":         Vec4{0.729, 0.333, 0.827, 1.0},
	"Mediumpurple":         Vec4{0.576, 0.439, 0.859, 1.0},
	"Mediumseagreen":       Vec4{0.235, 0.702, 0.443, 1.0},
	"Mediumslateblue":      Vec4{0.482, 0.408, 0.933, 1.0},
	"Mediumspringgreen":    Vec4{0.000, 0.980, 0.604, 1.0},
	"Mediumturquoise":      Vec4{0.282, 0.820, 0.800, 1.0},
	"Mediumvioletred":      Vec4{0.780, 0.082, 0.522, 1.0},
	"Midnightblue":         Vec4{0.098, 0.098, 0.439, 1.0},
	"Mintcream":            Vec4{0.961, 1.000, 0.980, 1.0},
	"Mistyrose":            Vec4{1.000, 0.894, 0.882, 1.0},
	"Moccasin":             Vec4{1.000, 0.894, 0.710, 1.0},
	"Navajowhite":          Vec4{1.000, 0.871, 0.678, 1.0},
	"Navy":                 Vec4{0.000, 0.000, 0.502, 1.0},
	"Oldlace":              Vec4{0.992, 0.961, 0.902, 1.0},
	"Olive":                Vec4{0.502, 0.502, 0.000, 1.0},
	"Olivedrab":            Vec4{0.420, 0.557, 0.137, 1.0},
	"Orange":               Vec4{1.000, 0.647, 0.000, 1.0},
	"Orangered":            Vec4{1.000, 0.271, 0.000, 1.0},
	"Orchid":               Vec4{0.855, 0.439, 0.839, 1.0},
	"Palegoldenrod":        Vec4{0.933, 0.910, 0.667, 1.0},
	"Palegreen":            Vec4{0.596, 0.984, 0.596, 1.0},
	"Paleturquoise":        Vec4{0.686, 0.933, 0.933, 1.0},
	"Palevioletred":        Vec4{0.859, 0.439, 0.576, 1.0},
	"Papayawhip":           Vec4{1.000, 0.937, 0.835, 1.0},
	"Peachpuff":            Vec4{1.000, 0.855, 0.725, 1.0},
	"Peru":                 Vec4{0.804, 0.522, 0.247, 1.0},
	"Pink":                 Vec4{1.000, 0.753, 0.796, 1.0},
	"Plum":                 Vec4{0.867, 0.627, 0.867, 1.0},
	"Powderblue":           Vec4{0.690, 0.878, 0.902, 1.0},
	"Purple":               Vec4{0.502, 0.000, 0.502, 1.0},
	"Red":                  Vec4{1.000, 0.000, 0.000, 1.0},
	"Rosybrown":            Vec4{0.737, 0.561, 0.561, 1.0},
	"Royalblue":            Vec4{0.255, 0.412, 0.882, 1.0},
	"Saddlebrown":          Vec4{0.545, 0.271, 0.075, 1.0},
	"Salmon":               Vec4{0.980, 0.502, 0.447, 1.0},
	"Sandybrown":           Vec4{0.957, 0.643, 0.376, 1.0},
	"Seagreen":             Vec4{0.180, 0.545, 0.341, 1.0},
	"Seashell":             Vec4{1.000, 0.961, 0.933, 1.0},
	"Sienna":               Vec4{0.627, 0.322, 0.176, 1.0},
	"Silver":               Vec4{0.753, 0.753, 0.753, 1.0},
	"Skyblue":              Vec4{0.529, 0.808, 0.922, 1.0},
	"Slateblue":            Vec4{0.416, 0.353, 0.804, 1.0},
	"Slategray":            Vec4{0.439, 0.502, 0.565, 1.0},
	"Slategrey":            Vec4{0.439, 0.502, 0.565, 1.0},
	"Snow":                 Vec4{1.000, 0.980, 0.980, 1.0},
	"Springgreen":          Vec4{0.000, 1.000, 0.498, 1.0},
	"Steelblue":            Vec4{0.275, 0.510, 0.706, 1.0},
	"Tan":                  Vec4{0.824, 0.706, 0.549, 1.0},
	"Teal":                 Vec4{0.000, 0.502, 0.502, 1.0},
	"Thistle":              Vec4{0.847, 0.749, 0.847, 1.0},
	"Tomato":               Vec4{1.000, 0.388, 0.278, 1.0},
	"Turquoise":            Vec4{0.251, 0.878, 0.816, 1.0},
	"Violet":               Vec4{0.933, 0.510, 0.933, 1.0},
	"Wheat":                Vec4{0.961, 0.871, 0.702, 1.0},
	"White":                Vec4{1.000, 1.000, 1.000, 1.0},
	"Whitesmoke":           Vec4{0.961, 0.961, 0.961, 1.0},
	"Yellow":               Vec4{1.000, 1.000, 0.000, 1.0},
	"Yellowgreen":          Vec4{0.604, 0.804, 0.196, 1.0},
}

func ColorName(cn string) Vec4 {

	return colorMap[cn]
}