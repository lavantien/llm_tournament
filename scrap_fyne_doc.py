import asyncio
from crawl4ai import AsyncWebCrawler, CacheMode


async def scrape_all_subpages(url, depth=3):
    visited_urls = set()  # To keep track of processed URLs
    results = []  # To store scraped content

    async with AsyncWebCrawler(verbose=True, depth=depth) as crawler:
        async def scrape_url(current_url):
            if current_url in visited_urls:
                return
            visited_urls.add(current_url)

            # Crawl the page
            result = await crawler.arun(
                url=current_url,
                exclude_external_links=True,
                process_iframes=True,
                remove_overlay_elements=True,
                cache_mode=CacheMode.ENABLED
            )
            results.append({"url": current_url, "content": result.markdown})

            # Extract links for further crawling
            for link in result.links:
                if link.startswith(url):  # Ensure the link is from the same domain
                    await scrape_url(link)

        await scrape_url(url)

    return results


async def main():
    urls = [
        "https://docs.fyne.io/",  # Add more URLs as needed
        "https://docs.fyne.io/started/",
        "https://docs.fyne.io/started/hello",
        "https://docs.fyne.io/started/demo",
        "https://docs.fyne.io/started/apprun",
        "https://docs.fyne.io/started/updating",
        "https://docs.fyne.io/started/windows",
        "https://docs.fyne.io/started/testing",
        "https://docs.fyne.io/started/packaging",
        "https://docs.fyne.io/started/webapp",
        "https://docs.fyne.io/started/metadata",
        "https://docs.fyne.io/started/cross-compiling",
        "https://docs.fyne.io/explore/canvas",
        "https://docs.fyne.io/explore/container",
        "https://docs.fyne.io/explore/widgets",
        "https://docs.fyne.io/explore/widgets",
        "https://docs.fyne.io/explore/dialogs",
        "https://docs.fyne.io/explore/icons",
        "https://docs.fyne.io/explore/shortcuts",
        "https://docs.fyne.io/explore/preferences",
        "https://docs.fyne.io/explore/translations",
        "https://docs.fyne.io/explore/systray",
        "https://docs.fyne.io/explore/binding",
        "https://docs.fyne.io/explore/compiling",
        "https://docs.fyne.io/canvas/rectangle",
        "https://docs.fyne.io/canvas/text",
        "https://docs.fyne.io/canvas/line",
        "https://docs.fyne.io/canvas/circle",
        "https://docs.fyne.io/canvas/image",
        "https://docs.fyne.io/canvas/raster",
        "https://docs.fyne.io/canvas/gradient",
        "https://docs.fyne.io/canvas/animation",
        "https://docs.fyne.io/container/box",
        "https://docs.fyne.io/container/grid",
        "https://docs.fyne.io/container/gridwrap",
        "https://docs.fyne.io/container/border",
        "https://docs.fyne.io/container/form",
        "https://docs.fyne.io/container/center",
        "https://docs.fyne.io/container/stack",
        "https://docs.fyne.io/container/apptabs",
        "https://docs.fyne.io/widget/label",
        "https://docs.fyne.io/widget/button",
        "https://docs.fyne.io/widget/entry",
        "https://docs.fyne.io/widget/choices",
        "https://docs.fyne.io/widget/form",
        "https://docs.fyne.io/widget/progressbar",
        "https://docs.fyne.io/widget/toolbar",
        "https://docs.fyne.io/collection/list",
        "https://docs.fyne.io/collection/table",
        "https://docs.fyne.io/collection/tree",
        "https://docs.fyne.io/binding/",
        "https://docs.fyne.io/binding/simple",
        "https://docs.fyne.io/binding/twoway",
        "https://docs.fyne.io/binding/conversion",
        "https://docs.fyne.io/binding/list",
        "https://docs.fyne.io/extend/custom-layout",
        "https://docs.fyne.io/extend/custom-widget",
        "https://docs.fyne.io/extend/bundle",
        "https://docs.fyne.io/extend/custom-theme",
        "https://docs.fyne.io/extend/extending-widgets",
        "https://docs.fyne.io/extend/numerical-entry",
        "https://docs.fyne.io/architecture/geometry",
        "https://docs.fyne.io/architecture/scaling",
        "https://docs.fyne.io/architecture/widgets",
        "https://docs.fyne.io/architecture/organisation",
        "https://docs.fyne.io/faq/layout",
        "https://docs.fyne.io/faq/theme",
        "https://docs.fyne.io/faq/troubleshoot",
        "https://docs.fyne.io/api/v2.5/upgrading.html",
    ]
    all_results = []

    for url in urls:
        data = await scrape_all_subpages(url)
        all_results.extend(data)

    # Write to a Markdown file
    with open("fyne_doc.md", "w", encoding="utf-8") as f:
        for result in all_results:
            f.write(f"# {result['url']}\n\n")
            f.write(result["content"])
            f.write("\n\n")

    print(f"Scraped {len(all_results)} pages across {
          len(urls)} base URLs successfully!")

if __name__ == "__main__":
    asyncio.run(main())
