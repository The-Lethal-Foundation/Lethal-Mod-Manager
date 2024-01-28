import React from 'react'
import type { FC } from 'react'

import {
  Card,
  CardContent,
  CardDescription,
  CardFooter,
  CardHeader,
  CardTitle,
} from '@/components/ui/card'
import { Button } from '@/components/ui/button'
import { Mod } from '@/types/mod'
import { Trash2Icon } from 'lucide-react'

interface ModCardProps {
  mod: Mod
  image?: string
}

export const ModCard: FC<ModCardProps> = ({
  mod,
  image = 'https://generated.vusercontent.net/placeholder.svg',
}) => {
  return (
    <Card>
      <CardHeader>
        <img
          src={image}
          alt="Mod #1"
          className="w-full h-48 object-cover rounded"
        />
      </CardHeader>
      <CardContent>
        <CardTitle>{mod.mod_name}</CardTitle>
        <CardDescription className="truncate">
          By {mod.mod_author}
        </CardDescription>
      </CardContent>
      <CardFooter>
        <Button variant="outline">
          <Trash2Icon className="h-4 w-4" />
        </Button>
      </CardFooter>
    </Card>
  )
}
