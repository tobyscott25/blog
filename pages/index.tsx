import type { NextPage } from 'next'
import Head from 'next/head'
import styles from '../styles/Home.module.css'


import { FontAwesomeIcon } from '@fortawesome/react-fontawesome'
import { IconLookup, IconDefinition, findIconDefinition } from '@fortawesome/fontawesome-svg-core'

const githubLookup: IconLookup = { prefix: 'fab', iconName: 'github' }
const githubIconDefinition: IconDefinition = findIconDefinition(githubLookup)

const twitterLookup: IconLookup = { prefix: 'fab', iconName: 'twitter' }
const twitterIconDefinition: IconDefinition = findIconDefinition(twitterLookup)

const ethereumLookup: IconLookup = { prefix: 'fab', iconName: 'ethereum' }
const ethereumIconDefinition: IconDefinition = findIconDefinition(ethereumLookup)


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
					<div className="text-6xl">{process.env.NEXT_PUBLIC_FULL_NAME}</div>
					<div className="text-2xl mt-4">Software Engineer, Greater Melbourne Area</div>
				</div>

				<a href={process.env.NEXT_PUBLIC_GITHUB_URL} className={styles.card} target="_blank" rel="noreferrer">
					<div className="text-xl">
						<FontAwesomeIcon icon={githubIconDefinition} className='mr-3' />
						GitHub &rarr;
					</div>
					<div className="text-lg">Projects, contributions &amp; dotfiles</div>
				</a>

				<a href={process.env.NEXT_PUBLIC_TWITTER_URL} className={styles.card} target="_blank" rel="noreferrer">
					<div className="text-xl">
					<FontAwesomeIcon icon={twitterIconDefinition} className='mr-3' />
						Twitter &rarr;
					</div>
					<div className="text-lg">Thoughts and opinions</div>
				</a>

				<div className='my-5 text-sm'>
					<FontAwesomeIcon icon={ethereumIconDefinition} className='mr-2' />
					<span className='font-mono'>0xF0B4cA8f326b7C14A7CC84355C057446587Dab0f</span>
				</div>
				
			</div>

		</div>
	)
}

export default Home