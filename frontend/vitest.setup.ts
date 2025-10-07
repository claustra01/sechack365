import "@testing-library/jest-dom/vitest";

if (typeof window !== "undefined" && window.matchMedia === undefined) {
	window.matchMedia = (query: string) => ({
		matches: false,
		media: query,
		onchange: null,
		addListener: () => {
			// no-op for Mantine color scheme detection in tests
		},
		removeListener: () => {
			// no-op for Mantine color scheme detection in tests
		},
		addEventListener: () => {
			// no-op for Mantine color scheme detection in tests
		},
		removeEventListener: () => {
			// no-op for Mantine color scheme detection in tests
		},
		dispatchEvent: () => false,
	});
}
