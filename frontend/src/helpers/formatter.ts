type PaymentStatus = 'pending' | 'completed' | 'processing' | 'failed'

export function formatRupiah(value: string | number): string {
  if (value === null || value === undefined || value === '')
    return '-'

  return new Intl.NumberFormat('id-ID', {
    minimumFractionDigits: 2,
    maximumFractionDigits: 2,
  }).format(Number(value))
}

const STATUS_MAP: Record<PaymentStatus, { color: string; label: string }> = {
  pending: {
    color: 'secondary',
    label: 'Pending',
  },
  completed: {
    color: 'success',
    label: 'Completed',
  },
  processing: {
    color: 'warning',
    label: 'Processing',
  },
  failed: {
    color: 'error',
    label: 'Failed',
  },
}

export function paymentStatusChip(status?: PaymentStatus) {
  return (
    STATUS_MAP[status as PaymentStatus] ?? {
      color: 'secondary',
      label: '-',
    }
  )
}

export function formatIndoTime(iso?: string) {
  if (!iso)
    return '-'

  const date = new Date(iso)

  return `${new Intl.DateTimeFormat('id-ID', {
    timeZone: 'Asia/Jakarta',
    day: '2-digit',
    month: 'short',
    year: 'numeric',
    hour: '2-digit',
    minute: '2-digit',
  }).format(date)} WIB`
}
