"use client";

import { createEntry } from "@/lib/api";
import { useRouter } from "next/navigation";
import { useState } from "react";

export default function Page() {
  const router = useRouter();
  const [text, setText] = useState("");
  const [source, setSource] = useState("open_journal");
  const [loading, setLoading] = useState(false);
  const [error, setError] = useState<string | null>(null);

  async function onSubmit() {
    setError(null);
    if (!text.trim()) {
      setError("Please enter some text.");
      return;
    }

    setLoading(true);
    try {
      const created = await createEntry({ text, source });
      router.push(`/entries/${encodeURIComponent(created.entry_id)}`);
    } catch (e: unknown) {
      const msg = e instanceof Error ? e.message : "Something went wrong";
      setError(msg);
    } finally {
      setLoading(false);
    }
  }

  return (
    <main className="min-h-screen bg-slate-950 text-slate-100">
      <div className="mx-auto max-w-3xl px-6 py-10">
        <h1 className="text-3xl font-semibold">PsychApp - Guided Journaling</h1>
        <p className="mt-2 text-slate-300">
          Write a short entry. You will receive a reflection and AQAL-style
          tags.
        </p>

        <div className="mt-8 rounded-2xl border border-slate-800 bg-slate-900 p-6 shadow">
          <label className="block text-sm text-slate-300">Entry type</label>
          <select
            className="mt-2 w-full rounded-lg border border-slate-700 bg-slate-950 px-3 py-2 text-slate-100"
            value={source}
            onChange={(e) => setSource(e.target.value)}
          >
            <option value="open_journal">Open journal</option>
            <option value="daily_checkin">Daily check-in</option>
            <option value="event_debrief">Event debrief</option>
          </select>

          <label className="mt-5 block text-sm text-slate-300">
            Your entry
          </label>
          <textarea
            className="mt-2 h-44 w-full resize-none rounded-lg border border-slate-700 bg-slate-950 px-3 py-2 text-slate-100"
            placeholder="What stood out today?"
            value={text}
            onChange={(e) => setText(e.target.value)}
          />

          {error && <p className="mt-3 text-sm text-red-300">{error}</p>}

          <button
            onClick={onSubmit}
            disabled={loading}
            className="mt-5 inline-flex items-center justify-center rounded-xl bg-cyan-700 px-4 py-2 font-medium hover:bg-cyan-600 disabled:cursor-not-allowed disabled:opacity-60"
          >
            {loading ? "Submitting..." : "Submit entry"}
          </button>
        </div>

        <div className="mt-6 text-sm text-slate-400">
          Tip: Keep it honest and concrete - thoughts, body sensations,
          relationships, and context.
        </div>
      </div>
    </main>
  );
}
