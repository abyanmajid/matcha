// for page navigation & to sort on leftbar

export type EachRoute = {
  title: string;
  href: string;
  noLink?: true; // noLink will create a route segment (section) but cannot be navigated
  items?: EachRoute[];
};

export const ROUTES: EachRoute[] = [
  {
    title: "Getting Started",
    href: "/getting-started",
    noLink: true,
    items: [
      { title: "Introduction", href: "/introduction" },
      {
        title: "Quickstart",
        href: "/quickstart",
      },
    ],
  },
  {
    title: "Guides",
    href: "/guides",
    noLink: true,
    items: [
      {
        title: "Routing and Middlewares",
        href: "/routing-and-middlewares",
      },
      {
        title: "OpenAPI",
        href: "/openapi",
      },
      {
        title: "Reference",
        href: "/reference",
      },
      {
        title: "JSON Web Token (JWT)",
        href: "/jwt",
      },
      {
        title: "Hashing",
        href: "/hashing",
      },
      {
        title: "Encryption",
        href: "/encryption",
      },
      {
        title: "Caching",
        href: "/caching",
      },
      {
        title: "Email",
        href: "/email",
      },
      {
        title: "Logging",
        href: "/logging",
      },
    ],
  },
  {
    title: "Other",
    href: "/other",
    noLink: true,
    items: [
      {
        title: "API Reference",
        href: "/reference",
      }
    ],
  }
];

type Page = { title: string; href: string };

function getRecurrsiveAllLinks(node: EachRoute) {
  const ans: Page[] = [];
  if (!node.noLink) {
    ans.push({ title: node.title, href: node.href });
  }
  node.items?.forEach((subNode) => {
    const temp = { ...subNode, href: `${node.href}${subNode.href}` };
    ans.push(...getRecurrsiveAllLinks(temp));
  });
  return ans;
}

export const page_routes = ROUTES.map((it) => getRecurrsiveAllLinks(it)).flat();
