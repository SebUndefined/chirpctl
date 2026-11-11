import type { ReactNode } from 'react';
import Link from '@docusaurus/Link';
import Layout from '@theme/Layout';
import Heading from '@theme/Heading';

import {
    ArrowRight,
    ArrowUpRight,
    Braces,
    Download,
    GitBranch,
    Layers3,
    Monitor,
    Network,
    RadioTower,
    ShieldCheck,
    Terminal,
    type LucideIcon,
} from 'lucide-react';
import { FaGithub } from 'react-icons/fa6';

import styles from './index.module.css';
import useBaseUrl from "@docusaurus/useBaseUrl";

const GITHUB_URL = 'https://github.com/sebundefined/chirpctl';
const RELEASES_URL = `${GITHUB_URL}/releases/latest`;

type Feature = {
    title: string;
    description: string;
    icon: LucideIcon;
};

const features: Feature[] = [
    {
        title: 'Native gRPC',
        description:
            'Interact directly with ChirpStack v4 through its gRPC API. No additional REST proxy required.',
        icon: RadioTower,
    },
    {
        title: 'Multiple contexts',
        description:
            'Keep your ChirpStack environments organized with reusable YAML-based connection profiles.',
        icon: Layers3,
    },
    {
        title: 'Secure connections',
        description:
            'Configure TLS, custom CA certificates, server names and API tokens for your infrastructure.',
        icon: ShieldCheck,
    },
    {
        title: 'Tenant-aware',
        description:
            'Use tenant-specific defaults and command options to work with the right LoRaWAN resources.',
        icon: Network,
    },
    {
        title: 'Automation-ready',
        description:
            'Choose table, JSON, YAML or CSV output to integrate chirpctl into scripts and CI/CD pipelines.',
        icon: Braces,
    },
    {
        title: 'Cross-platform',
        description:
            'Run native Go binaries on Linux, macOS and Windows, with AMD64 and ARM64 builds.',
        icon: Monitor,
    },
];

const commands = [
    {
        comment: 'Configure your ChirpStack connection',
        command: 'chirpctl config init',
    },
    {
        comment: 'Check the active context',
        command: 'chirpctl config current-context',
    },
    {
        comment: 'List gateways as JSON',
        command: 'chirpctl get gateways --output json',
    },
];

function HeroSection(): ReactNode {
    const bannerUrl = useBaseUrl('/img/chirpctl-banner.png');
    return (
        <section className={styles.hero}>
            <div className={styles.container}>
                <div className={styles.heroGrid}>
                    <div className={styles.heroContent}>
                        <div className={styles.eyebrow}>
                            <span className={styles.statusIndicator} />
                            <span>CHIRPSTACK V4</span>
                            <span className={styles.eyebrowSeparator}>/</span>
                            <span>GRPC</span>
                            <span className={styles.eyebrowSeparator}>/</span>
                            <span>GO</span>
                        </div>

                        <Heading as="h1" className={styles.heroTitle}>
                            ChirpStack.
                            <span>At your fingertips.</span>
                        </Heading>

                        <p className={styles.heroDescription}>
                            Manage your LoRaWAN infrastructure from the command line.
                            A cross-platform CLI built for operators, developers and
                            automation.
                        </p>

                        <div className={styles.heroActions}>
                            <Link to="/docs/" className={styles.primaryButton}>
                                Get started
                                <ArrowRight size={18} aria-hidden="true" />
                            </Link>

                            <Link href={GITHUB_URL} className={styles.secondaryButton}>
                                <FaGithub size={18} aria-hidden="true" />
                                View on GitHub
                            </Link>
                        </div>

                        <div className={styles.platformList}>
                            <span className={styles.platformLabel}>AVAILABLE FOR</span>
                            <span>Linux</span>
                            <span className={styles.platformDot}>·</span>
                            <span>macOS</span>
                            <span className={styles.platformDot}>·</span>
                            <span>Windows</span>
                        </div>
                    </div>

                    <div className={styles.heroVisual}>
                        <div className={styles.heroImageFrame}>
                            <img
                                src={bannerUrl}
                                alt="Go Gopher technician installing a ChirpStack gateway on a radio tower"
                                className={styles.heroImage}
                                loading="eager"
                                fetchPriority="high"
                            />
                        </div>

                        <div className={styles.heroImageCaption}>
                            <RadioTower size={15} aria-hidden="true" />
                            <span>Your LoRaWAN infrastructure, one command away.</span>
                        </div>
                    </div>
                </div>
            </div>

            <div className={styles.heroBottomGlow} aria-hidden="true" />
        </section>
    );
}

function SectionHeading({
                            eyebrow,
                            title,
                            description,
                        }: {
    eyebrow: string;
    title: string;
    description: string;
}): ReactNode {
    return (
        <div className={styles.sectionHeading}>
            <span className={styles.sectionEyebrow}>{eyebrow}</span>

            <Heading as="h2" className={styles.sectionTitle}>
                {title}
            </Heading>

            <p className={styles.sectionDescription}>{description}</p>
        </div>
    );
}

function FeatureCard({
                         feature,
                         index,
                     }: {
    feature: Feature;
    index: number;
}): ReactNode {
    const Icon = feature.icon;

    return (
        <article className={styles.featureCard}>
            <div className={styles.featureCardTop}>
                <div className={styles.featureIcon}>
                    <Icon size={24} strokeWidth={1.8} aria-hidden="true" />
                </div>

                <span className={styles.featureIndex}>
          {String(index + 1).padStart(2, '0')}
        </span>
            </div>

            <Heading as="h3" className={styles.featureTitle}>
                {feature.title}
            </Heading>

            <p className={styles.featureDescription}>
                {feature.description}
            </p>
        </article>
    );
}

function FeaturesSection(): ReactNode {
    return (
        <section
            className={styles.featuresSection}
            aria-labelledby="features-heading"
        >
            <div className={styles.container}>
                <SectionHeading
                    eyebrow="BUILT FOR THE TERMINAL"
                    title="Everything you need. Nothing in your way."
                    description="A focused command-line experience for everyday ChirpStack operations, with the flexibility to fit into your existing workflows."
                />

                <div className={styles.featureGrid} id="features-heading">
                    {features.map((feature, index) => (
                        <FeatureCard
                            key={feature.title}
                            feature={feature}
                            index={index}
                        />
                    ))}
                </div>
            </div>
        </section>
    );
}

function TerminalWindow(): ReactNode {
    return (
        <div className={styles.terminalWindow}>
            <div className={styles.terminalHeader}>
                <div className={styles.terminalControls} aria-hidden="true">
                    <span className={styles.terminalRed} />
                    <span className={styles.terminalYellow} />
                    <span className={styles.terminalGreen} />
                </div>

                <div className={styles.terminalTitle}>
                    <Terminal size={15} aria-hidden="true" />
                    <span>chirpctl — terminal</span>
                </div>

                <span className={styles.terminalLanguage}>SHELL</span>
            </div>

            <div className={styles.terminalBody}>
                {commands.map((item) => (
                    <div className={styles.commandBlock} key={item.command}>
                        <div className={styles.commandComment}>
                            # {item.comment}
                        </div>

                        <div className={styles.commandLine}>
                            <span className={styles.commandPrompt}>$</span>
                            <code>{item.command}</code>
                        </div>
                    </div>
                ))}

                <div className={styles.terminalCursorLine}>
                    <span className={styles.commandPrompt}>$</span>
                    <span className={styles.terminalCursor} aria-hidden="true" />
                </div>
            </div>

            <div className={styles.terminalFooter}>
        <span className={styles.terminalReady}>
          <span className={styles.readyDot} />
          Ready for automation
        </span>

                <span>chirpctl</span>
            </div>
        </div>
    );
}

function WorkflowSection(): ReactNode {
    return (
        <section className={styles.workflowSection}>
            <div className={styles.container}>
                <div className={styles.workflowGrid}>
                    <div className={styles.workflowContent}>
            <span className={styles.sectionEyebrow}>
              SIMPLE BY DESIGN
            </span>

                        <Heading as="h2" className={styles.sectionTitle}>
                            A familiar workflow.
                            <span className={styles.titleAccent}>
                {' '}A more powerful terminal.
              </span>
                        </Heading>

                        <p className={styles.sectionDescription}>
                            Configure your connection once, select your environment
                            and interact with ChirpStack directly from your shell.
                        </p>

                        <div className={styles.workflowSteps}>
                            <div className={styles.workflowStep}>
                                <div className={styles.stepIcon}>
                                    <ShieldCheck size={20} aria-hidden="true" />
                                </div>

                                <div>
                                    <h3>Configure securely</h3>
                                    <p>
                                        Set up connection profiles with TLS and API
                                        authentication.
                                    </p>
                                </div>
                            </div>

                            <div className={styles.workflowStep}>
                                <div className={styles.stepIcon}>
                                    <GitBranch size={20} aria-hidden="true" />
                                </div>

                                <div>
                                    <h3>Work across environments</h3>
                                    <p>
                                        Manage reusable contexts without constantly
                                        re-entering connection settings.
                                    </p>
                                </div>
                            </div>

                            <div className={styles.workflowStep}>
                                <div className={styles.stepIcon}>
                                    <Braces size={20} aria-hidden="true" />
                                </div>

                                <div>
                                    <h3>Automate with structured output</h3>
                                    <p>
                                        Integrate commands into scripts using JSON,
                                        YAML or CSV.
                                    </p>
                                </div>
                            </div>
                        </div>
                    </div>

                    <div className={styles.workflowVisual}>
                        <TerminalWindow />

                        <div className={styles.outputFormats}>
                            <span>OUTPUT FORMATS</span>
                            <div className={styles.formatBadges}>
                                <span>TABLE</span>
                                <span>JSON</span>
                                <span>YAML</span>
                                <span>CSV</span>
                            </div>
                        </div>
                    </div>
                </div>
            </div>
        </section>
    );
}

function CompatibilitySection(): ReactNode {
    return (
        <section className={styles.compatibilitySection}>
            <div className={styles.container}>
                <div className={styles.compatibilityCard}>
                    <div className={styles.compatibilityIcon}>
                        <RadioTower size={25} strokeWidth={1.7} aria-hidden="true" />
                    </div>

                    <div className={styles.compatibilityText}>
            <span className={styles.sectionEyebrow}>
              BUILT FOR CHIRPSTACK
            </span>

                        <Heading as="h2">
                            Native v4 support. Designed to evolve.
                        </Heading>

                        <p>
                            chirpctl currently targets the ChirpStack v4 gRPC API,
                            with an architecture designed to accommodate additional
                            API versions over time.
                        </p>
                    </div>

                    <Link
                        href={GITHUB_URL}
                        className={styles.compatibilityLink}
                    >
                        Explore the project
                        <ArrowUpRight size={18} aria-hidden="true" />
                    </Link>
                </div>
            </div>
        </section>
    );
}

function GetStartedSection(): ReactNode {
    return (
        <section className={styles.ctaSection}>
            <div className={styles.container}>
                <div className={styles.ctaCard}>
                    <div className={styles.ctaGlow} aria-hidden="true" />

                    <div className={styles.ctaContent}>
                        <div className={styles.ctaIcon}>
                            <Terminal size={27} aria-hidden="true" />
                        </div>

                        <Heading as="h2">
                            Your LoRaWAN infrastructure.
                            <span> One CLI.</span>
                        </Heading>

                        <p>
                            Download chirpctl, configure your first context,
                            and start working with ChirpStack from your terminal.
                        </p>

                        <div className={styles.ctaActions}>
                            <Link to="/docs/" className={styles.primaryButton}>
                                Read the documentation
                                <ArrowRight size={18} aria-hidden="true" />
                            </Link>

                            <Link
                                href={RELEASES_URL}
                                className={styles.secondaryButton}
                            >
                                <Download size={18} aria-hidden="true" />
                                Download latest release
                            </Link>
                        </div>
                    </div>
                </div>
            </div>
        </section>
    );
}

export default function Home(): ReactNode {
    return (
        <Layout
            title="ChirpStack CLI for LoRaWAN infrastructure"
            description="chirpctl is a cross-platform command-line tool for managing ChirpStack v4 through gRPC. Secure contexts, structured output and automation-ready workflows."
        >
            <main className={styles.page}>
                <HeroSection />
                <FeaturesSection />
                <WorkflowSection />
                <CompatibilitySection />
                <GetStartedSection />
            </main>
        </Layout>
    );
}
