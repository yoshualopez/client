import {namedConnect} from '../../util/container'
import * as Types from '../../constants/types/chat2'
import * as Constants from '../../constants/chat2'
import * as Styles from '../../styles'
import ChooseConversation from './choose-conversation'

type OwnProps = {
  dropdownButtonStyle?: Styles.StylesCrossPlatform | null
  filter?: string
  onSelect: (conversationIDKey: Types.ConversationIDKey) => void
  onSetFilter?: (filter: string) => void
  selected: Types.ConversationIDKey
}

const mapStateToProps = (state, ownProps) => ({
  _conv: state.chat2.metaMap.get(ownProps.selected),
})

const mapDispatchToProps = () => ({})

const mergeProps = (stateProps, _, ownProps: OwnProps) => ({
  ...ownProps,
  selectedText: stateProps._conv
    ? Constants.getConversationLabel(stateProps._conv, false)
    : 'Choose a conversation',
})

export default namedConnect(mapStateToProps, mapDispatchToProps, mergeProps, 'ChooseConversation')(
  ChooseConversation
)
