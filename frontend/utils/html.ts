import DOMPurify from 'dompurify';
import parse from 'html-react-parser';

const parseUrlLink = (text: string) => {
  const regex = /((https?):\/\/[\w!?/+\-_~;.,*&@#$%()'[\]]+)/g;
  return text.replace(regex, (url) => `<a href="${url}" target="_blank">${url}</a>`);
}

export const parseHtml = (html: string) => {
  html = html.replaceAll("\n", "<br>");
  const sanitized =  DOMPurify.sanitize(html, {
    ALLOWED_TAGS: ["p", "span", "br"],
  });
  return parse(parseUrlLink(sanitized));
}
