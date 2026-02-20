export interface CodeReviewRequest {
  code: string
  language: string
  context?: string
}

export interface CodeReviewResponse {
  score: number
  suggestions: string[]
  issues: Issue[]
  summary: string
  review_id: string
}

export interface Issue {
  type: 'error' | 'warning' | 'info'
  severity: 'high' | 'medium' | 'low'
  line: number
  message: string
  suggestion: string
}
