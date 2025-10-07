import type { ReactNode } from "react";

import { MantineProvider } from "@mantine/core";
import { render } from "@testing-library/react";

/**
 * Mantineのテーマコンテキストを含めてコンポーネントを描画する。
 */
export const renderWithMantine = (ui: ReactNode) => {
	return render(<MantineProvider>{ui}</MantineProvider>);
};
