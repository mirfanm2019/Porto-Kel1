import { Check, Sparkle } from "lucide-react";
import { profile } from "../../data/profile.js";
import { api } from "../../lib/api.js";
import { useFetch } from "../../hooks/useFetch.js";
import SectionHeading from "../ui/SectionHeading.jsx";
import Container from "../ui/Container.jsx";
import Decor from "../ui/Decor.jsx";
import { ContactIcon, TechLogo, techName } from "../ui/Icons.jsx";

function Label({children}){
    return <h3 className="font-mono text-label uppercase tracking-label text-stone">{children}</h3>
}

export default function About(){
    const { data: me } = useFetch(api.getProfile);
    const fullName = me?.fullName || profile.fullName;
    const location = me?.location || profile.location;
    const photo = me?.photoUrl || profile.photo;
    const bio = me?.bio || profile.pitch;
    const about = me?.about || profile.about;

    const contacts = profile.contacts.map((contact) => {
        if (contact.type === "email" && me?.email) {
            return { ...contact, label: me.email, href: `mailto:${me.email}` };
        }
        if (contact.type === "location" && me?.location) {
            return { ...contact, label: me.location };
        }
        return contact;
    });
    return(
        <section id="profil" aria-labelledby="profil-title" className="relative overflow-hidden bg-paper bg-grid py-24 md:py-36 lg:slide">

            <div className="absolute top-20 left-0 hidden h-[420px] w-[220px] bg-dots fade-edges md:block" aria-hidden="true" />
            <div className="absolute top-0 right-0 h-[520px] w-[520px] rounded-full bg-pink/10 blur-3xl" aria-hidden="true" />

            <Container className="relative">
                <div className="grid items-end gap-10 lg:grid-cols-[1fr_400px] lg:gap-6 xl:grid-cols-[1fr_480px]">
                    <SectionHeading
                        id="profil-title"
                        eyebrow="Profil kelompok"
                        dot="bg-sage"
                        line1="Hii, get to know us!"
                        line2="siapakah kami?"
                    />
                    <div className="rounded-3xl border border-white bg-white/60 p-6 shadow-card backdrop-blur lg:p-5">
                        <p className="leading-loose text-ink lg:leading-relaxed">{bio}</p>
                        <p className="mt-4 font-serif text-xl text-violet-deep italic lg:mt-2">↳ {profile.motto}</p>
                    </div>
                </div>
                <div className="mt-20 grid gap-12 md:grid-cols-[300px_1fr] lg:mt-slide lg:grid-cols-[240px_1fr_220px] lg:gap-10 xl:grid-cols-[280px_1fr_280px]">
                    <figure className="relative mx-auto w-full max-w-[340px] -rotate-2 self-start rounded-[22px] bg-cream p-3 shadow-float">
                        <Decor src="/images/mask.png" className="-top-6 -right-6 w-20 rotate-6" />
                        <img
                            src={photo}
                            alt={`Foto ${fullName}`}
                            className="aspect-[3/4] w-full rounded-[14px] object-cover object-center"
                        />
                        <figcaption className="flex items-end justify-between px-3 pt-5 pb-3 lg:pt-3 lg:pb-1">
                            <div>
                                <p className="font-serif text-xl italic">{fullName}</p>
                                <p className="mt-1 flex items-center gap-2 font-mono text-[10px] uppercase tracking-label text-violet">
                                    <span className="size-1.5 rounded-full bg-pink" aria-hidden="true" />
                                    {location}
                                </p>
                            </div>
                            <Sparkle className="size-6 fill-violet text-violet" aria-hidden="true" />
                        </figcaption>
                    </figure>

                    <div>
                        <Label>Tentang</Label>
                        <p className="mt-4 max-w-[46ch] text-lg leading-[1.7] lg:mt-3 lg:text-base lg:leading-relaxed">{about}</p>

                        <ul className="mt-7 space-y-3 lg:mt-4 lg:space-y-2">
                            {profile.highlights.map((item) => (
                                <li key={item} className="flex items-center gap-3 text-[15px]">
                                    <Check className="size-4 shrink-0 text-sage" aria-hidden="true" />
                                    {item}
                                </li>
                            ))}
                        </ul>

                        <div className="mt-10 lg:mt-5">
                            <Label>Kontak</Label>
                            <ul className="mt-4 space-y-3 lg:mt-3 lg:grid lg:grid-cols-2 lg:gap-x-4 lg:gap-y-2 lg:space-y-0">
                                {contacts.map((contact) => {
                                    const content = (
                                        <>
                                            <ContactIcon type={contact.type} className="size-5 shrink-0 text-muted" />
                                            {contact.label}
                                        </>
                                    );
                                    return (
                                        <li key={contact.type}>
                                            {contact.href ? (
                                                <a href={contact.href} className="flex items-center gap-3 hover:text-violet">
                                                    {content}
                                                </a>
                                            ) : (
                                                <span className="flex items-center gap-3">{content}</span>
                                            )}
                                        </li>
                                    );
                                })}
                            </ul>
                        </div>
                    </div>

                    <div className="md:col-span-2 lg:col-span-1">
                        <Label>Teknologi</Label>
                        <ul className="mt-5 flex flex-wrap gap-6 lg:mt-3 lg:gap-5">
                            {me?.tech?.length > 0
                                ? me.tech.map((t) => (
                                    <li key={t.id} title={t.name} className="drop-shadow-[0_6px_10px_rgb(74_31_168/0.15)]">
                                        {t.iconUrl ? (
                                            <img src={t.iconUrl} alt="" className="size-8" />
                                        ) : (
                                            <span className="text-sm font-medium">{t.name}</span>
                                        )}
                                        <span className="sr-only">{t.name}</span>
                                    </li>
                                ))
                                : profile.tech.map((id) => (
                                    <li key={id} title={techName(id)} className="drop-shadow-[0_6px_10px_rgb(74_31_168/0.15)]">
                                        <TechLogo id={id} className="size-8" />
                                        <span className="sr-only">{techName(id)}</span>
                                    </li>
                                ))}
                        </ul>

                        <div className="mt-10 lg:mt-5">
                            <Label>Keahlian</Label>
                            <ul className="mt-4 space-y-3 pl-3 text-[15px] lg:mt-3 lg:space-y-2">
                                {profile.skills.map((skill) => (
                                    <li key={skill}>{skill}</li>
                                ))}
                            </ul>
                        </div>
                    </div>
                </div>

                {profile.teamMembers.length > 0 && (
                    <section aria-labelledby="team-title" className="mt-16 lg:mt-12">
                        <h3 id="team-title" className="font-mono text-label uppercase tracking-label text-stone">
                            Teman Satu Tim
                        </h3>
                        <div className="mt-8 space-y-12">
                            {profile.teamMembers.map((member) => (
                                <article key={member.id} className="grid items-start gap-12 md:grid-cols-[300px_1fr] lg:grid-cols-[240px_1fr] lg:gap-10 xl:grid-cols-[280px_1fr]">
                                    <figure className="relative mx-auto w-full max-w-[340px] -rotate-2 self-start rounded-[22px] bg-cream p-3 shadow-float">
                                        <Decor src="/images/mask.png" className="-top-6 -right-6 w-20 rotate-6" />
                                        {member.photo ? (
                                            <img
                                                src={member.photo}
                                                alt={`Foto ${member.name}`}
                                                className="aspect-[3/4] w-full rounded-[14px] object-cover object-center"
                                            />
                                        ) : (
                                            <div
                                                role="img"
                                                aria-label={`Foto ${member.name} belum ditambahkan`}
                                                className="grid aspect-[3/4] place-items-center rounded-[14px] bg-lilac text-sm text-muted"
                                            >
                                                Foto belum ditambahkan
                                            </div>
                                        )}
                                        <figcaption className="flex items-end justify-between px-3 pt-5 pb-3 lg:pt-3 lg:pb-1">
                                            <div>
                                                <p className="font-serif text-xl italic">{member.name}</p>
                                                <p className="mt-1 flex items-center gap-2 font-mono text-[10px] uppercase tracking-label text-violet">
                                                    <span className="size-1.5 rounded-full bg-pink" aria-hidden="true" />
                                                    {member.role}
                                                </p>
                                            </div>
                                            <Sparkle className="size-6 fill-violet text-violet" aria-hidden="true" />
                                        </figcaption>
                                    </figure>

                                    <div>
                                        <Label>Tentang</Label>
                                        <p className="mt-4 max-w-[46ch] text-lg leading-[1.7] lg:mt-3 lg:text-base lg:leading-relaxed">
                                            {member.about}
                                        </p>

                                        <ul className="mt-7 space-y-3 lg:mt-4 lg:space-y-2">
                                            {member.highlights.map((highlight) => (
                                                <li key={highlight} className="flex items-center gap-3 text-[15px]">
                                                    <Check className="size-4 shrink-0 text-sage" aria-hidden="true" />
                                                    {highlight}
                                                </li>
                                            ))}
                                        </ul>

                                        <div className="mt-10 lg:mt-5">
                                            <Label>Kontak</Label>
                                            <ul className="mt-4 space-y-3 lg:mt-3 lg:grid lg:grid-cols-2 lg:gap-x-4 lg:gap-y-2 lg:space-y-0">
                                                {member.contacts.map((contact) => {
                                                    const content = (
                                                        <>
                                                            <ContactIcon type={contact.type} className="size-5 shrink-0 text-muted" />
                                                            {contact.label}
                                                        </>
                                                    );
                                                    return (
                                                        <li key={contact.type}>
                                                            {contact.href ? (
                                                                <a href={contact.href} className="flex items-center gap-3 hover:text-violet">
                                                                    {content}
                                                                </a>
                                                            ) : (
                                                                <span className="flex items-center gap-3">{content}</span>
                                                            )}
                                                        </li>
                                                    );
                                                })}
                                            </ul>
                                        </div>
                                    </div>
                                </article>
                            ))}
                        </div>
                    </section>
                )}
            </Container>
        </section>
    )
}