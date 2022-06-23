import styled from 'styled-components'

export const ButtonText = styled.button`
  font-size: ${props => props.theme.fontSize.small};
  font-weight: bold;
  line-height: 19px;
  letter-spacing: 1.34px;
  text-transform: uppercase;
`

export const Body1 = styled.p`
  font-size: ${props => props.theme.fontSize.base};
  line-height: 24px;
  letter-spacing: 0.44px;
`

export const Body2 = styled.p`
  font-size: ${props => props.theme.fontSize.small};
  line-height: 22px;
  letter-spacing: 0.25px;
`

export const Caption = styled.span`
  font-size: ${props => props.theme.fontSize.tiny};
  line-height: 16px;
  letter-spacing: 0.4px;
`

export const H1 = styled.h1`
  font-size: ${props => props.theme.fontSize.header};
  line-height: 145px;
  font-weight: 300;
  letter-spacing: -1.54px;
`

export const H2 = styled.h2`
  font-size: ${props => props.theme.fontSize['3xlarge']};
  line-height: 90px;
  font-weight: 300;
  letter-spacing: -0.48px;
`

export const H3 = styled.h3`
  font-size: ${props => props.theme.fontSize['2xlarge']};
  line-height: 72px;
`

export const H4 = styled.h4`
  font-size: ${props => props.theme.fontSize.xlarge};
  line-height: 51px;
  letter-spacing: 0.26px;
`

export const H5 = styled.h5`
  font-size: ${props => props.theme.fontSize.large};
  line-height: 35px;
`

export const H6 = styled.h6`
  font-size: ${props => props.theme.fontSize.medium};
  font-weight: 600;
  line-height: 30px;
  letter-spacing: 0.15px;
`

export const Overline = styled.span`
  font-size: ${props => props.theme.fontSize.micro};
  line-height: 16px;
  letter-spacing: 1.5px;
  text-transform: uppercase;
`

export const Subtitle1 = styled.p`
  font-size: ${props => props.theme.fontSize.base};
  font-weight: bold;
  line-height: 25px;
  letter-spacing: 0.14px;
`

export const Subtitle2 = styled.p`
  font-size: ${props => props.theme.fontSize.small};
  font-weight: 600;
  line-height: 24px;
  letter-spacing: 0.1px;
`
