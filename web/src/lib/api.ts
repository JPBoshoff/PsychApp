export type CreateEntryRequest = {
  text: string;
  source?: string;
  metadata?: Record<string, string>;
};

export type EntryAnalysis = {
  quadrant_distribution?: Record<string, number>;
  themes?: string[];
  mirror_reflection?: {
    summary?: string;
    clarifying_questions?: string[];
  };
  safety?: {
    risk?: string;
    recommended_action?: string;
    signals?: string[];
  };
};

export type CreateEntryResponse = {
  entry_id: string;
  created_at: string;
  analysis: EntryAnalysis;
  request_id?: string;
  mock_notice?: string;
};

export type GetEntryResponse = {
  entry_id: string;
  created_at: string;
  text: string;
  source?: string;
  metadata?: Record<string, string>;
  analysis: EntryAnalysis;
};

const baseUrl = process.env.NEXT_PUBLIC_API_BASE_URL;

function assertBaseUrl() {
  if (!baseUrl) throw new Error("NEXT_PUBLIC_API_BASE_URL is not set");
  return baseUrl;
}

export async function createEntry(
  payload: CreateEntryRequest,
): Promise<CreateEntryResponse> {
  const res = await fetch(`${assertBaseUrl()}/entries`, {
    method: "POST",
    headers: { "Content-Type": "application/json" },
    body: JSON.stringify(payload),
  });

  if (!res.ok) {
    const text = await res.text();
    throw new Error(`Create entry failed: ${res.status} ${text}`);
  }

  return res.json();
}

export async function getEntry(entryId: string): Promise<GetEntryResponse> {
  const res = await fetch(
    `${assertBaseUrl()}/entries/${encodeURIComponent(entryId)}`,
    {
      method: "GET",
    },
  );

  if (!res.ok) {
    const text = await res.text();
    throw new Error(`Get entry failed: ${res.status} ${text}`);
  }

  return res.json();
}
