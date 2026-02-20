import axios from 'axios'
import { CodeReviewRequest, CodeReviewResponse } from '../types'

const API_URL = import.meta.env.VITE_API_URL || 'http://localhost:8080'

const api = axios.create({
  baseURL: `${API_URL}/api/v1`,
  headers: {
    'Content-Type': 'application/json',
  },
})

export const reviewCode = async (
  code: string,
  language: string
): Promise<CodeReviewResponse> => {
  const request: CodeReviewRequest = {
    code,
    language,
  }

  const response = await api.post<CodeReviewResponse>('/review', request)
  return response.data
}

export const healthCheck = async (): Promise<{ status: string }> => {
  const response = await api.get('/health')
  return response.data
}
