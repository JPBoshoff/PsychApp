"use client";

import { EntryListItem, listEntries } from "@/lib/api";
import Link from "next/link";
import { useEffect, useState } from "react";

export default function TimelinePage() {
  const [items, setItems] = useState<EntryListItem[]>([]);
  const [error, setError] = useState<string | null>(null);

  useEffect(() => {
    (async () => {
      try {
        const res = await listEntries(50);
        setItems(res.items ?? []);
      } catch (e: unknown) {
        const msg = e instanceof Error ? e.message : "Failed to load timeline";
        setError(msg);
      }
    })();
  }, []);

  return (
    <main className="min-h-screen bg-slate-950 text-slate-100">
      <div className="mx-auto max-w-3xl px-6 py-10">
        <div className="flex items-center justify-between">
          <div>
            <h1 className="text-2xl font-semibold">Timeline</h1>
            <p className="mt-1 text-slate-300">Your recent journal entries</p>
          </div>
          <Link href="/" className="text-sm text-cyan-400 hover:text-cyan-300">
            + New entry
          </Link>
        </div>

        {error && <p className="mt-6 text-sm text-red-300">{error}</p>}

        {!error && items.length === 0 && (
          <div className="mt-6 rounded-2xl border border-slate-800 bg-slate-900 p-6 text-slate-300">
            No entries yet. Create your first one.
          </div>
        )}

        <div className="mt-6 space-y-3">
          {items.map((it) => (
            <Link
              key={it.entry_id}
              href={`/entries/${encodeURIComponent(it.entry_id)}`}
              className="block rounded-2xl border border-slate-800 bg-slate-900 p-5 hover:border-slate-700"
            >
              <div className="flex items-start justify-between gap-4">
                <div className="min-w-0">
                  <div className="text-sm text-slate-400">
                    {new Date(it.created_at).toLocaleString()}
                    {it.source ? <span> - {it.source}</span> : null}
                  </div>
                  <div className="mt-2 line-clamp-2 text-slate-100">
                    {it.excerpt}
                  </div>

                  {it.themes?.length ? (
                    <div className="mt-3 flex flex-wrap gap-2">
                      {it.themes.slice(0, 6).map((t) => (
                        <span
                          key={t}
                          className="rounded-full bg-slate-800 px-3 py-1 text-xs text-slate-200"
                        >
                          {t}
                        </span>
                      ))}
                    </div>
                  ) : null}
                </div>

                <div className="shrink-0 text-slate-500">→</div>
              </div>
            </Link>
          ))}
        </div>
      </div>
    </main>
  );
}
