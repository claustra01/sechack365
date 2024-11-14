"use client";

import { postApiV1AuthLogout } from "@/openapi";

export default function LogoutPage() {
  postApiV1AuthLogout().then((response) => {
    if (response.status === 200) {
      // FIXME: routerを使うようにする
      window.location.href = "/";
    }
  }).catch((error) => {
    console.error(error);
    alert("ログアウトに失敗しました");
  });

  return <></>
}