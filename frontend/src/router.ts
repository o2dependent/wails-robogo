import { writable } from "svelte/store";

/* NOTE: issue with access the file system with import.meta.glob to generate Screens list dynamically
 * you must also add your screen route in the Router ifelse statement and add the screen to the Screens array
 * if this can be fixed that'd be awesome, but if it works it works
 */
export const Screens = ["Home", "About", "Bounce", "NotFound"];

export const route = writable<{ name: (typeof Screens)[number]; data?: any }>({
	name: "Home",
});

export const goto = (name: (typeof Screens)[number], data?: any) =>
	route.update(() => {
		if (Screens.includes(name)) {
			return { name, data };
		}
		return { name: "", data: undefined };
	});
