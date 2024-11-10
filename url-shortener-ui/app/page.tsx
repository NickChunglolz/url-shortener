"use client";

import { useState } from "react";
import { Button } from "@/components/ui/button";
import { Input } from "@/components/ui/input";
import { toast } from "sonner";
import { Link2, Copy, ExternalLink } from "lucide-react";
import { Card } from "@/components/ui/card";

export default function Home() {
  const [url, setUrl] = useState("");
  const [shortUrl, setShortUrl] = useState("");
  const [loading, setLoading] = useState(false);

  const isValidUrl = (urlString: string) => {
    try {
      new URL(urlString);
      return true;
    } catch (err) {
      return false;
    }
  };

  const shortenUrl = async () => {
    if (!isValidUrl(url)) {
      toast.error("Please enter a valid URL");
      return;
    }

    setLoading(true);
    try {
      const response = await fetch("https://api.shrtco.de/v2/shorten?url=" + encodeURIComponent(url));
      const data = await response.json();
      if (data.ok) {
        setShortUrl(data.result.full_short_link);
        toast.success("URL shortened successfully!");
      } else {
        toast.error("Failed to shorten URL");
      }
    } catch (error) {
      toast.error("An error occurred");
    } finally {
      setLoading(false);
    }
  };

  const copyToClipboard = async () => {
    try {
      await navigator.clipboard.writeText(shortUrl);
      toast.success("Copied to clipboard!");
    } catch (err) {
      toast.error("Failed to copy");
    }
  };

  return (
    <main className="min-h-screen bg-gradient-to-b from-neutral-50 to-neutral-100 dark:from-neutral-950 dark:to-neutral-900">
      <div className="container mx-auto px-4 py-16 max-w-3xl">
        <div className="text-center mb-12">
          <div className="inline-block p-3 bg-primary/10 rounded-full mb-4">
            <Link2 className="w-8 h-8 text-primary" />
          </div>
          <h1 className="text-4xl font-bold mb-4">URL Shortener</h1>
          <p className="text-muted-foreground">
            Transform your long URLs into short, shareable links in seconds
          </p>
        </div>

        <Card className="p-6 shadow-lg">
          <div className="flex flex-col md:flex-row gap-4">
            <Input
              type="url"
              placeholder="Enter your long URL here..."
              value={url}
              onChange={(e) => setUrl(e.target.value)}
              className="flex-1"
            />
            <Button
              onClick={shortenUrl}
              disabled={loading}
              className="md:w-32"
            >
              {loading ? "Shortening..." : "Shorten"}
            </Button>
          </div>

          {shortUrl && (
            <div className="mt-6 p-4 bg-muted/50 rounded-lg">
              <div className="flex items-center justify-between gap-4">
                <div className="flex-1 truncate">
                  <a
                    href={shortUrl}
                    target="_blank"
                    rel="noopener noreferrer"
                    className="text-primary hover:underline flex items-center gap-2"
                  >
                    {shortUrl} <ExternalLink className="w-4 h-4" />
                  </a>
                </div>
                <Button
                  variant="secondary"
                  size="icon"
                  onClick={copyToClipboard}
                  className="shrink-0"
                >
                  <Copy className="w-4 h-4" />
                </Button>
              </div>
            </div>
          )}
        </Card>

        <div className="mt-12 grid grid-cols-1 md:grid-cols-3 gap-6">
          <div className="text-center p-6">
            <h3 className="font-semibold mb-2">Fast & Reliable</h3>
            <p className="text-sm text-muted-foreground">
              Instant URL shortening with high availability
            </p>
          </div>
          <div className="text-center p-6">
            <h3 className="font-semibold mb-2">Secure</h3>
            <p className="text-sm text-muted-foreground">
              Your data is safe and URLs are protected
            </p>
          </div>
          <div className="text-center p-6">
            <h3 className="font-semibold mb-2">Simple</h3>
            <p className="text-sm text-muted-foreground">
              Easy to use interface with no registration required
            </p>
          </div>
        </div>
      </div>
    </main>
  );
}