"use client";

import { getEntry, GetEntryResponse } from "@/lib/api";
import Link from "next/link";
import { useParams } from "next/navigation";
import { useEffect, useMemo, useState } from "react";

function QuadrantBar({ label, value }: { label: string; value: number }) {
  const pct = Math.round((value || 0) * 100);
  return (
    <div className="flex items-center gap-3">
      <div className="w-10 text-sm text-slate-300">{label}</div>
      <div className="h-2 flex-1 rounded-full bg-slate-800">
        <div
          className="h-2 rounded-full bg-cyan-700"
          style={{ width: `${pct}%` }}
        />
      </div>
      <div className="w-12 text-right text-sm text-slate-300">{pct}%</div>
    </div>
  );
}

export default function EntryPage() {
  const params = useParams<{ id: string }>();
  const id = params?.id;
  const [data, setData] = useState<GetEntryResponse | null>(null);
  const [error, setError] = useState<string | null>(null);

  useEffect(() => {
    if (!id) return;
    (async () => {
      try {
        const entry = await getEntry(id);
        setData(entry);
      } catch (e: unknown) {
        const msg = e instanceof Error ? e.message : "Failed to load entry";
        setError(msg);
      }
    })();
  }, [id]);

  const q = useMemo(() => data?.analysis?.quadrant_distribution ?? {}, [data]);
  const themes = data?.analysis?.themes ?? [];
  const reflection = data?.analysis?.mirror_reflection;

  return (
    <main className="min-h-screen bg-slate-950 text-slate-100">
      <div className="mx-auto max-w-3xl px-6 py-10">
        <Link href="/" className="text-sm text-cyan-400 hover:text-cyan-300">
          ← New entry
        </Link>

        <h1 className="mt-3 text-2xl font-semibold">Entry</h1>

        {error && <p className="mt-4 text-sm text-red-300">{error}</p>}
        {!data && !error && <p className="mt-4 text-slate-300">Loading...</p>}

        {data && (
          <div className="mt-6 space-y-6">
            <div className="rounded-2xl border border-slate-800 bg-slate-900 p-6">
              <div className="text-sm text-slate-400">Created</div>
              <div className="mt-1">
                {new Date(data.created_at).toLocaleString()}
              </div>

              <div className="mt-5 text-sm text-slate-400">Text</div>
              <div className="mt-1 whitespace-pre-wrap text-slate-100">
                {data.text}
              </div>
            </div>

            <div className="rounded-2xl border border-slate-800 bg-slate-900 p-6">
              <h2 className="text-lg font-semibold">Reflection</h2>
              <p className="mt-3 text-slate-200">
                {reflection?.summary ?? "—"}
              </p>

              {reflection?.clarifying_questions?.length ? (
                <>
                  <div className="mt-5 text-sm font-medium text-slate-200">
                    Questions
                  </div>
                  <ul className="mt-2 list-disc space-y-1 pl-5 text-slate-200">
                    {reflection.clarifying_questions.map((qq, i) => (
                      <li key={i}>{qq}</li>
                    ))}
                  </ul>
                </>
              ) : null}
            </div>

            <div className="rounded-2xl border border-slate-800 bg-slate-900 p-6">
              <h2 className="text-lg font-semibold">Quadrants</h2>
              <div className="mt-4 space-y-3">
                <QuadrantBar label="UL" value={q["UL"] ?? 0} />
                <QuadrantBar label="UR" value={q["UR"] ?? 0} />
                <QuadrantBar label="LL" value={q["LL"] ?? 0} />
                <QuadrantBar label="LR" value={q["LR"] ?? 0} />
              </div>
            </div>

            <div className="rounded-2xl border border-slate-800 bg-slate-900 p-6">
              <h2 className="text-lg font-semibold">Themes</h2>
              <div className="mt-3 flex flex-wrap gap-2">
                {themes.length ? (
                  themes.map((t) => (
                    <span
                      key={t}
                      className="rounded-full bg-slate-800 px-3 py-1 text-sm text-slate-200"
                    >
                      {t}
                    </span>
                  ))
                ) : (
                  <span className="text-slate-400">—</span>
                )}
              </div>
            </div>
          </div>
        )}
      </div>
    </main>
  );
}
