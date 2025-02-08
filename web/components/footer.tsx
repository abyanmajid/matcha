import Link from "next/link";
import { buttonVariants } from "./ui/button";
import { Github } from "lucide-react";

export function Footer() {
  return (
    <footer className="border-t w-full h-16">
      <div className="container flex items-center sm:justify-between justify-center sm:gap-0 gap-4 h-full text-muted-foreground text-sm flex-wrap sm:py-0 py-3 max-sm:px-4">
        <div className="flex items-center gap-3">
          <p className="text-center">
            Built by{" "}
            <Link
              className="px-1 underline underline-offset-2"
              href="https://github.com/abyanmajid"
            >
              abyanmajid
            </Link>
          </p>
        </div>

        <div className="gap-4 items-center hidden md:flex">
          <FooterButtons />
        </div>
      </div>
    </footer>
  );
}

export function FooterButtons() {
  return (
    <>
      <Link
        href="https://github.com/abyanmajid/matcha"
        className={buttonVariants({ variant: "outline", size: "sm" })}
      >
        <Github className="h-4 w-4 mr-2 fill-current" />
        GitHub
      </Link>
    </>
  );
}
