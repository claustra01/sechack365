"use client";

import axios from "axios";

export default function AxiosProvider({
	children,
}: {
	children: React.ReactNode;
}) {
	axios.defaults.baseURL = process.env.NEXT_PUBLIC_HOST;
	return children;
}
