import type { NextPage } from 'next'
import Head from 'next/head'
import styles from '../styles/Home.module.css'


import { FontAwesomeIcon } from '@fortawesome/react-fontawesome'
import { IconLookup, IconDefinition, findIconDefinition } from '@fortawesome/fontawesome-svg-core'

const githubLookup: IconLookup = { prefix: 'fab', iconName: 'github' }
const githubIconDefinition: IconDefinition = findIconDefinition(githubLookup)

const twitterLookup: IconLookup = { prefix: 'fab', iconName: 'twitter' }
const twitterIconDefinition: IconDefinition = findIconDefinition(twitterLookup)


const Home: NextPage = () => {
	return (
		<div className={styles.container}>
			<Head>
				<title>Toby Scott</title>
				<meta name="description" content="Software Engineer, Greater Melbourne Area" />
				<meta name="viewport" content="width=device-width, initial-scale=1.0" />
				<link rel="icon" href="/favicon.ico" />
			</Head>

			<div className={styles.main}>

				<div className="my-5 text-center">
					<div className="text-6xl">Toby Scott</div>
					<div className="text-2xl mt-4">Software Engineer, Greater Melbourne Area</div>
				</div>

				<a href="https://github.com/tobyscott25" className={styles.card} target="_blank">
					<div className="text-xl">
						<FontAwesomeIcon icon={githubIconDefinition} className='mr-3' />
						GitHub &rarr;
					</div>
					<div className="text-lg">Projects, contributions &amp; dotfiles</div>
				</a>

				<a href="https://twitter.com/tobyscott25" className={styles.card} target="_blank">
					<div className="text-xl">
					<FontAwesomeIcon icon={twitterIconDefinition} className='mr-3' />
						Twitter &rarr;
					</div>
					<div className="text-lg">Thoughts and opinions</div>
				</a>
				
			</div>

		</div>
	)
}

export default Home