import '../styles/globals.css'
import type { AppProps } from 'next/app'

import { library, config } from '@fortawesome/fontawesome-svg-core'
import { fas } from '@fortawesome/free-solid-svg-icons'
import { fab } from '@fortawesome/free-brands-svg-icons'
import { far } from '@fortawesome/free-regular-svg-icons'

import '@fortawesome/fontawesome-svg-core/styles.css'

library.add(fas, fab, far)
config.autoAddCss = false


function MyApp({ Component, pageProps }: AppProps) {
	return <Component {...pageProps} />
}

export default MyApp
