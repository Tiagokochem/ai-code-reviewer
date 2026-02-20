import { useState } from 'react'
import { reviewCode } from '../services/api'
import { CodeReviewResponse } from '../types'
import './CodeReviewer.css'

const CodeReviewer = () => {
  const [code, setCode] = useState('')
  const [language, setLanguage] = useState('javascript')
  const [loading, setLoading] = useState(false)
  const [review, setReview] = useState<CodeReviewResponse | null>(null)
  const [error, setError] = useState<string | null>(null)

  const languages = [
    { value: 'javascript', label: 'JavaScript' },
    { value: 'typescript', label: 'TypeScript' },
    { value: 'vue', label: 'Vue.js' },
    { value: 'php', label: 'PHP' },
    { value: 'laravel', label: 'Laravel' },
    { value: 'go', label: 'Go' },
    { value: 'python', label: 'Python' },
    { value: 'java', label: 'Java' },
    { value: 'rust', label: 'Rust' },
    { value: 'cpp', label: 'C++' },
    { value: 'c', label: 'C' },
  ]

  const handleSubmit = async (e: React.FormEvent) => {
    e.preventDefault()
    setError(null)
    setReview(null)

    if (!code.trim()) {
      setError('Por favor, cole algum código para análise')
      return
    }

    setLoading(true)

    try {
      const result = await reviewCode(code, language)
      setReview(result)
    } catch (err) {
      setError(err instanceof Error ? err.message : 'Erro ao analisar código')
    } finally {
      setLoading(false)
    }
  }

  const getScoreColor = (score: number) => {
    if (score >= 80) return '#4caf50'
    if (score >= 60) return '#ff9800'
    return '#f44336'
  }

  return (
    <div className="code-reviewer">
      <form onSubmit={handleSubmit} className="review-form">
        <div className="form-group">
          <label htmlFor="language">Linguagem:</label>
          <select
            id="language"
            value={language}
            onChange={(e) => setLanguage(e.target.value)}
            className="language-select"
          >
            {languages.map((lang) => (
              <option key={lang.value} value={lang.value}>
                {lang.label}
              </option>
            ))}
          </select>
        </div>

        <div className="form-group">
          <label htmlFor="code">Cole seu código:</label>
          <textarea
            id="code"
            value={code}
            onChange={(e) => setCode(e.target.value)}
            placeholder="Cole ou digite o código aqui..."
            className="code-input"
            rows={15}
          />
        </div>

        <button type="submit" disabled={loading} className="submit-btn">
          {loading ? 'Analisando...' : '🔍 Analisar Código'}
        </button>
      </form>

      {error && (
        <div className="error-message">
          <strong>❌ Erro:</strong> {error}
        </div>
      )}

      {review && (
        <div className="review-results">
          <div className="score-section">
            <h2>Score de Qualidade</h2>
            <div
              className="score-circle"
              style={{ borderColor: getScoreColor(review.score) }}
            >
              <span className="score-value">{review.score}</span>
              <span className="score-max">/100</span>
            </div>
          </div>

          <div className="summary-section">
            <h3>📝 Resumo</h3>
            <p>{review.summary}</p>
          </div>

          {review.suggestions.length > 0 && (
            <div className="suggestions-section">
              <h3>💡 Sugestões</h3>
              <ul>
                {review.suggestions.map((suggestion, idx) => (
                  <li key={idx}>{suggestion}</li>
                ))}
              </ul>
            </div>
          )}

          {review.issues.length > 0 && (
            <div className="issues-section">
              <h3>⚠️ Problemas Encontrados</h3>
              <div className="issues-list">
                {review.issues.map((issue, idx) => (
                  <div key={idx} className={`issue issue-${issue.severity}`}>
                    <div className="issue-header">
                      <span className="issue-type">{issue.type}</span>
                      <span className="issue-severity">{issue.severity}</span>
                      {issue.line > 0 && (
                        <span className="issue-line">Linha {issue.line}</span>
                      )}
                    </div>
                    <p className="issue-message">{issue.message}</p>
                    {issue.suggestion && (
                      <p className="issue-suggestion">
                        💡 {issue.suggestion}
                      </p>
                    )}
                  </div>
                ))}
              </div>
            </div>
          )}

          <div className="review-id">
            <small>Review ID: {review.review_id}</small>
          </div>
        </div>
      )}
    </div>
  )
}

export default CodeReviewer
