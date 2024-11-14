import { postApiV1Posts } from "@/openapi";
import { Newpost } from "@/openapi/schemas";
import { Box, Button, Textarea } from "@mantine/core";
import { useState } from "react";

export const NewPost = () => {
  const [text, setText] = useState<string>("");

  const handleChange = (e: React.ChangeEvent<HTMLTextAreaElement>) => {
    setText(e.currentTarget.value);
  }

  const handleSubmit = () => {
    if (text.length === 0) {
      return;
    }
    const newPost: Newpost = {
      content: text,
    };
    postApiV1Posts(newPost).then((response) => {
        // FIXME: routerを使うようにする
        window.location.href = "/";
    }).catch((error) => {
      alert(error);
    })
  }    

  return (
    <Box style={{display: "flex", flexDirection: "column", padding: "24px", gap: "24px"}}>
      <Textarea
        placeholder="Your content here..."
        autosize
        minRows={4}
        onChange={handleChange}
      />
      <Button 
        style={{alignSelf: "flex-end"}}
        onClick={handleSubmit}
      >
        New Post
      </Button>
    </Box>
  );
}
