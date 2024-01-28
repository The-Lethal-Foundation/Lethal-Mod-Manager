import React from 'react'
import { cn } from '@/lib/utils'

export interface ISVGProps extends React.SVGProps<SVGSVGElement> {
  size?: number
  className?: string
  theme?: 'white' | 'black' // Add theme prop
}

export const LoadingSpinner = ({
  size = 24,
  className,
  theme,
  ...props
}: ISVGProps) => {
  // Determine stroke color based on theme
  const strokeColor =
    theme === 'white' ? '#FFF' : theme === 'black' ? '#000' : 'currentColor'

  return (
    <svg
      xmlns="http://www.w3.org/2000/svg"
      width={size}
      height={size}
      {...props}
      viewBox="0 0 24 24"
      fill="none"
      stroke={strokeColor} // Apply stroke color
      strokeWidth="2"
      strokeLinecap="round"
      strokeLinejoin="round"
      className={cn('animate-spin', className)}
    >
      <path d="M21 12a9 9 0 1 1-6.219-8.56" />
    </svg>
  )
}
