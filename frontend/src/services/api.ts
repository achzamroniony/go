// Centralized API client for communicating with the Go Fiber backend.

const API_BASE_URL = import.meta.env.VITE_API_URL || 'http://localhost:8080';

export interface APIResponse<T = any> {
  success: boolean;
  message?: string;
  data?: T;
  errors?: any;
}

export interface HelloData {
  message: string;
  topics: string[];
  version: string;
}

/**
 * Fetch utility wrapper that sets headers and resolves JSON.
 */
async function apiFetch<T>(endpoint: string, options?: RequestInit): Promise<APIResponse<T>> {
  const url = `${API_BASE_URL}${endpoint}`;
  
  try {
    const response = await fetch(url, {
      headers: {
        'Content-Type': 'application/json',
        ...(options?.headers || {}),
      },
      ...options,
    });
    
    if (!response.ok) {
      const errorText = await response.text();
      let errorData;
      try {
        errorData = JSON.parse(errorText);
      } catch {
        errorData = { message: errorText || 'Terjadi kesalahan server' };
      }
      return {
        success: false,
        message: errorData.message || `HTTP Error: ${response.status}`,
        errors: errorData.errors || null,
      };
    }
    
    return await response.json();
  } catch (error: any) {
    return {
      success: false,
      message: error.message || 'Koneksi ke backend gagal. Pastikan API backend menyala.',
    };
  }
}

/**
 * Check backend service health status.
 */
export async function checkHealth(): Promise<APIResponse> {
  return apiFetch('/api/health');
}

/**
 * Fetch hello message and topics from the backend.
 */
export async function fetchHello(): Promise<APIResponse<HelloData>> {
  return apiFetch<HelloData>('/api/hello');
}
