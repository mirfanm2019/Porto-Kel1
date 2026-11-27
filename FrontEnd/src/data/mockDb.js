import {
    mockJourney,
    mockProjects,
    mockCertificates,
    mockHero,
    mockProfile,
    mockFooter,
    mockTechnologies,
} from "./mock.js";

const STORAGE_KEY = "portfolio-mock-db-v3";
const OLD_HERO_DESCRIPTION =
    "Saya membantu tim meluncurkan produk yang kuat dan mudah dikelola — memadukan pemikiran produk, kejelasan visual, dan pemecahan masalah berbasis kode.";
const PREVIOUS_HERO_DESCRIPTION =
    "Kami membantu tim meluncurkan produk yang kuat dan mudah dikelola — memadukan pemikiran produk, kejelasan visual, dan pemecahan masalah berbasis kode.";
const OLD_PROFILE_ABOUT =
    "Saya bermitra dengan tim untuk mengubah ide produk menjadi pengalaman yang dipoles dan mudah dikelola. Pendekatan saya menyeimbangkan kejelasan visual, kelayakan teknis, dan kolaborasi yang efektif.";
const OLD_MOCK_PROFILE_ABOUT =
    "Kami bermitra untuk mengubah ide produk menjadi pengalaman yang dipoles dan mudah dikelola. Pendekatan kami menyeimbangkan kejelasan visual, kelayakan teknis, dan kolaborasi yang efektif.";

function seed() {
    return structuredClone({
        hero: mockHero,
        profile: mockProfile,
        footer: mockFooter,
        journey: mockJourney,
        projects: mockProjects,
        certificates: mockCertificates,
        technologies: mockTechnologies,
    });
}

let memory = null;

export function readDb() {
    if (memory) return memory;
    try {
        const saved = localStorage.getItem(STORAGE_KEY);
        memory = saved ? JSON.parse(saved) : seed();
        let migrated = false;
        if (
            memory.certificates?.length === 3 &&
            memory.certificates.every((certificate) => certificate.title === "[Judul Sertifikasi]" && !certificate.credentialId)
        ) {
            memory.certificates = structuredClone(mockCertificates);
            migrated = true;
        }
        if (memory.hero?.description === OLD_HERO_DESCRIPTION) {
            memory.hero.description = mockHero.description;
            migrated = true;
        }
        if (memory.hero?.displayName === "[Nama Full]" && memory.hero?.headline === "Bangun dan luncurkan.") {
            memory.hero.displayName = mockHero.displayName;
            memory.hero.roleLabel = mockHero.roleLabel;
            memory.hero.headline = mockHero.headline;
            if (!memory.hero.heroText) memory.hero.heroText = mockHero.heroText;
            if (!memory.hero.supportText) memory.hero.supportText = mockHero.supportText;
            if (memory.hero.extraText === "Akademik × PTI") memory.hero.extraText = mockHero.extraText;
            if (memory.hero.ctaLabel === "Jelajahi proyek") memory.hero.ctaLabel = mockHero.ctaLabel;
            if (memory.hero.ctaUrl === "#proyek") memory.hero.ctaUrl = mockHero.ctaUrl;
            migrated = true;
        }
        if (memory.hero?.description === PREVIOUS_HERO_DESCRIPTION) {
            memory.hero.description = mockHero.description;
            migrated = true;
        }
        if (memory.profile?.about === OLD_PROFILE_ABOUT) {
            memory.profile.about = mockProfile.about;
            migrated = true;
        }
        if (
            memory.profile?.photoUrl === "/images/profile.jpg" &&
            memory.profile?.fullName === "[Nama Full]" &&
            memory.profile?.role === "Fullstack Developer" &&
            memory.profile?.location === "[Kota, Provinsi]" &&
            memory.profile?.email === "email@contoh.com" &&
            memory.profile?.bio === "Sedikit niat di setiap interaksi." &&
            memory.profile?.about === OLD_MOCK_PROFILE_ABOUT
        ) {
            memory.profile = structuredClone(mockProfile);
            migrated = true;
        }
        if (migrated) {
            writeDb(memory);
        }
    } catch {
        memory = seed();
    }
    return memory;
}

export function writeDb(db) {
    memory = db;
    try {
        localStorage.setItem(STORAGE_KEY, JSON.stringify(db));
    } catch {
        // Penuh (misalnya karena gambar besar) atau diblokir: biarkan di memori saja.
    }
}

export function resetDb() {
    try {
        localStorage.removeItem(STORAGE_KEY);
    } catch {
        // abaikan
    }
    memory = seed();
    return memory;
}
