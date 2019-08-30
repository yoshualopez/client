import * as I from 'immutable'
import * as Types from './types/deeplinks'

export const linkIsKeybaseLink = (link: string) => link.startsWith('keybase://')

export const linkFromConvAndMessage = (conv: string, messageID: number) =>
  `keybase://chat/${conv}/${messageID}`

export const makeState = I.Record<Types._State>({
  keybaseLinkError: '',
})
