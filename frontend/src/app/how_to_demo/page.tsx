"use client";
import { Header } from "@/components/Header/Header";
import { PageTemplate } from "@/components/Template/PageTemplate";
import { Box, ScrollArea } from "@mantine/core";
import { IconArrowBackUp } from "@tabler/icons-react";
import Markdown from "react-markdown";

export default function HowToDemoPage() {
	const content = `
# TL;DR
このSNSはActivityPubとNostr双方のプロトコルに対応しています。

ユーザーを作成して適当なリモートアカウントをフォローし、リモートアカウントから投稿することでこのサーバーのユーザーのタイムラインに表示されます。
また、リモートユーザーからこのアカウントをフォローし、このアカウントから投稿することで、リモートユーザーのタイムラインに表示されます。

# Caution
このデモサーバーはテスト用のため、データは定期的に削除されます。
連合SNSの特性上、このサーバーとフォロー関係を作成するActivityPubやNostrのアカウントは使い捨てのものを推奨します。

以下のインスタンスで使い捨てのアカウントが発行できます。リモートアカウントを作成する場合はこれらのインスタンスの使用を強く推奨します。
- ActivityPub: [ActivityPub Academy](https://activitypub.academy/)
- Nostr: [Nostter](https://nostter.app/)

# Search/Follow
フォローを行うには、ユーザーを連合から検索する必要があります。

ユーザーIDはActivityPubの場合\`@user@domain\`の形、Nostrの場合\`npub1xxxxxx....\`の形をしています。
このサーバーのユーザーのIDはログイン後のメニューから飛ぶことができるプロフィールページにどちらも表示されています。
リモートインスタンス上でこれらを検索することで、このユーザーをリモートから検索・フォローすることが出来ます。

リモートのアカウントをこのサーバーで検索する場合、検索フォームにユーザーIDを入力して検索してください。
もし検索フォームがまだ実装されていない/見つからない場合は\`https://yojohan-demo.claustra01.net/profile/{ユーザーID}\`を開いてください。
そこからそのリモートユーザーのフォローや投稿閲覧が可能です。

# Article
このサーバーのユーザーは、このサーバー上で記事を投稿することができます。記事はMarkdown形式で記述することができます。
PNG/JPEG形式の画像を添付することも可能です。

記事を投稿するとURLを含んだオブジェクトがローカル・リモートのタイムラインに流れます。ここにリプライをするという形でコメントを残すことができます。
  `;

	return (
		<PageTemplate>
			<Header title={"How to Demo"} icon={<IconArrowBackUp />} />
			<ScrollArea h={"calc( 100vh - 78px )"}>
				<Box p={24}>
					<Markdown>{content}</Markdown>
				</Box>
			</ScrollArea>
		</PageTemplate>
	);
}
