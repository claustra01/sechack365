import DOMPurify from 'dompurify';
import parse from 'html-react-parser';

const parseUrlLink = (text: string) => {
  const regex = /((https?):\/\/[^\s]+)/g;
  return text.replace(regex, (url) => `<a href="${url}" target="_blank">${url}</a>`);
}

export const parseHtml = (html: string) => {
  const sanitized =  DOMPurify.sanitize(html, {
    ALLOWED_TAGS: ["p", "span", "br"],
  });
  return parse(parseUrlLink(sanitized));
}
