import type { NextPage } from 'next'
import Head from 'next/head'
import styles from '../styles/Home.module.css'

const Home: NextPage = () => {
	return (
		<div className={styles.container}>
			<Head>
				<title>Toby Scott</title>
				<meta name="description" content="About Toby Scott" />
				<link rel="icon" href="/favicon.ico" />
			</Head>

			<div className={styles.main}>

				<div className="my-5 text-center">
					<div className="text-6xl">Toby Scott</div>
					<div className="text-2xl mt-4">Software Engineer, Greater Melbourne Area</div>
				</div>

				<a href="https://github.com/tobyscott25" className={styles.card}>
					<div className="text-xl">GitHub &rarr;</div>
					<div className="text-lg">See my contributions, projects, dotfiles, etc.</div>
				</a>

				<a href="https://twitter.com/tobyscott25" className={styles.card}>
					<div className="text-xl">Twitter &rarr;</div>
					<div className="text-lg">Read the thoughts and opinions that I share.</div>
				</a>
				
			</div>

		</div>
	)
}

export default Home